#!/bin/bash

# SURICATA_DIR_HOST needs to be set in the env!

source .env 

mkdir -p ${SURICATA_DIR_HOST}/{etc,lib/rules,log}

docker-compose -f docker-compose-suricata.yml up
