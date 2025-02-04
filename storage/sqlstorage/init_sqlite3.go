package sqlstorage

import "github.com/jmoiron/sqlx"

const (
	dbSchema = `CREATE TABLE IF NOT EXISTS ips (
		id INTEGER NOT NULL PRIMARY KEY,
		ip TEXT NOT NULL
	);
	CREATE UNIQUE INDEX IF NOT EXISTS ips_ip ON ips (ip);
`
)

func initSqlite3(dbx *sqlx.DB) error {

	_, err := dbx.Exec(dbSchema)

	return err
}
