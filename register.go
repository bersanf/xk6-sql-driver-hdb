// Package ramsql contains SAP HANA driver registration for xk6-sql.
package hdb

import (
	"github.com/grafana/xk6-sql/sql"

	// Blank import required for initialization of driver.
	_ "github.com/SAP/go-hdb/driver"
)

func init() {
	sql.RegisterModule("hdb")
}
