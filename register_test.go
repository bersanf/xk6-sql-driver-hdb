package hdb

import (
	_ "embed"
	"os"
	"testing"

	"github.com/grafana/xk6-sql/sqltest"
)

//go:embed testdata/script.js
var script string

// dsnEnv contains the name of the environment variable holding the SAP HANA
// connection string used by the integration test. The test is skipped when
// unset, because it requires a reachable SAP HANA instance.
const dsnEnv = "K6_SQL_HDB_DSN"

func TestIntegration(t *testing.T) { //nolint:paralleltest
	dsn, found := os.LookupEnv(dsnEnv)
	if !found {
		t.Skipf("set %s to an hdb:// connection string to run this test", dsnEnv)
	}

	sqltest.RunScript(t, "hdb", dsn, script)
}
