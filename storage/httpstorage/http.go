package httpstorage

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/diegohce/droneip/storage"
)

type httpStorage struct {
	url string
}

func openHttpStorage(dsn string) (storage.Storager, error) {

	fs := httpStorage{
		url: dsn,
	}
	return &fs, nil
}

func (s *httpStorage) Save(ip string) error {

	body := fmt.Sprintf(`{"ip": "%s"}`, ip)

	res, err := http.Post(s.url, "application/json", strings.NewReader(body))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, _ = io.ReadAll(res.Body)

	return nil
}

func (s *httpStorage) List() ([]string, error) {
	return nil, errors.New("no local copy of banned ips")
}

func (s *httpStorage) Close() error {
	return nil
}

func init() {
	storage.Register("http", openHttpStorage)
}
