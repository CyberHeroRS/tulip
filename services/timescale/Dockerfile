FROM timescale/timescaledb:latest-pg15

RUN apk add build-base make clang19 llvm19 git
COPY timescale/tulip /tulip
RUN cd /tulip && make USE_PGXS=1 install && cd / && \
	git clone https://github.com/ossc-db/pg_hint_plan --branch PG15 && \
	cd pg_hint_plan && make USE_PGXS=1 install

COPY schema/system.sql /docker-entrypoint-initdb.d/100_system.sql
COPY schema/functions.sql /docker-entrypoint-initdb.d/101_functions.sql
COPY schema/schema.sql /docker-entrypoint-initdb.d/102_schema.sql
