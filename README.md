# xk6-sql-driver-hdb

Database driver extension for [xk6-sql](https://github.com/grafana/xk6-sql) k6 extension to support SAP HANA database.

## Example

```JavaScript file=examples/example.js
import sql from "k6/x/sql";
import driver from "k6/x/sql/driver/hdb";

const db = sql.open(driver, "hdb://HANA_USER:HANA_PASSWORD@HANA_HOST:30015");

export function setup() {
  db.exec(`
  CREATE COLUMN TABLE test_table (
    a INT GENERATED BY DEFAULT AS IDENTITY,
    b TEXT);
  `);
}

export function teardown() {
  db.exec(`
    DROP TABLE test_table;
  `);
  db.close();
}

export default function () {
  let result = db.exec(`
    INSERT INTO test_table (b) VALUES ('One');
  `);
  console.log(`${result.rowsAffected()} rows inserted`);

  let rows = db.query("SELECT a,b FROM test_table WHERE b = $1", 'One');
  for (const row of rows) {
    console.log(`key: ${row.a}, value: ${row.b}`);
  }
}
```

## Usage

Check the [xk6-sql documentation](https://github.com/grafana/xk6-sql) on how to use this database driver.
