all: test build

test: *.go testdata/*.js
	go test -count 1 ./...

build: k6

k6: *.go go.mod go.sum
	xk6 build --with github.com/grafana/xk6-sql@latest --with github.com/bersanf/xk6-sql-driver-hdb=.

# Requires a reachable SAP HANA instance, eg.
#   K6_SQL_HDB_DSN='hdb://myUser:myPassword@localhost:30015' make example
example: k6
	./k6 run -e K6_SQL_HDB_DSN=$(K6_SQL_HDB_DSN) examples/example.js

.PHONY: all test build example
