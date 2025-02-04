package sqlstorage

import (
	"net/url"
	"strings"

	"github.com/diegohce/droneip/logger"
	"github.com/diegohce/droneip/storage"
	"github.com/jmoiron/sqlx"
)

type sqlStorage struct {
	dbx *sqlx.DB
}

func openSQLStorage(dsn string) (storage.Storager, error) {

	u, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	if u.Scheme == "sqlite" {
		dsn = strings.Replace(dsn, "sqlite", "file", 1)
	}

	dbx, err := sqlx.Open(u.Scheme, dsn)
	if err != nil {
		return nil, err
	}

	if err = InitDB(u.Scheme, dbx); err != nil {
		dbx.Close()
		return nil, err
	}

	s := sqlStorage{
		dbx: dbx,
	}

	return &s, nil
}

func (s *sqlStorage) Save(ip string) error {
	return saveIP(s.dbx, ip)
}

func (s *sqlStorage) List() ([]string, error) {
	return listIPs(s.dbx)
}

func init() {
	storage.Register("sql", openSQLStorage)
}

func saveIP(dbx *sqlx.DB, ip string) error {

	_, err := dbx.NamedExec("INSERT INTO ips (ip) VALUES (:ip)", map[string]interface{}{"ip": ip})
	logger.LogInfo("saveIP: error", "err", err.Error()).Write()
	return nil
}

func listIPs(dbx *sqlx.DB) ([]string, error) {
	var ips []string

	err := dbx.Select(&ips, "SELECT ip FROM ips ORDER BY ip")

	return ips, err
}
