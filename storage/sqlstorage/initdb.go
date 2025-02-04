package sqlstorage

import "github.com/jmoiron/sqlx"

type InitDBFunc func(*sqlx.DB) error

var (
	initdbfuncs = map[string]InitDBFunc{
		"sqlite": initSqlite3,
		"mysql":  initMySQL,
	}
)

func InitDB(name string, dbx *sqlx.DB) error {

	fn, ok := initdbfuncs[name]
	if !ok {
		return nil
	}

	return fn(dbx)
}
