package main

import (
	"go-importer/internal/pkg/db"
	"reflect"
	"sync"
	"time"

	"github.com/cloudflare/ahocorasick"
)

type flagIDMatcher struct {
	contents []string
	matcher  *ahocorasick.Matcher
}

type flagIDCache struct {
	mu      sync.Mutex
	updated time.Time
	current *flagIDMatcher
}

func (cache *flagIDCache) get(database *db.Database, lifetime int, interval time.Duration) (*flagIDMatcher, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if cache.current != nil && time.Since(cache.updated) < interval {
		return cache.current, nil
	}

	contents, err := database.FlagIdsQuery(lifetime)
	if err != nil {
		return nil, err
	}
	if cache.current == nil || !reflect.DeepEqual(cache.current.contents, contents) {
		cache.current = &flagIDMatcher{contents: contents}
		if len(contents) > 0 {
			cache.current.matcher = ahocorasick.NewStringMatcher(contents)
		}
	}
	cache.updated = time.Now()
	return cache.current, nil
}
