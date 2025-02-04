package sqlstorage

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

var mysqlCreateDB = []string{
	`CREATE TABLE IF NOT EXISTS ips (
		id INTEGER NOT NULL AUTO_INCREMENT,
		ip TEXT NOT NULL,
		PRIMARY KEY (id)
	)`,
	`ALTER TABLE ips ADD UNIQUE INDEX ips_ip (ip(128))`,
}

func initMySQL(dbx *sqlx.DB) error {

	for _, q := range mysqlCreateDB {
		_, err := dbx.Exec(q)
		if !strings.Contains(q, "INDEX") && err != nil {
			return err
		}
	}

	return nil
}
