package filestorage

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/diegohce/droneip/storage"
)

type fileStorage struct {
	filePath string
	mutex    sync.RWMutex
}

func openFileStorage(dsn string) (storage.Storager, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	fs := fileStorage{
		filePath: u.Path,
	}
	return &fs, nil
}

func (s *fileStorage) Save(ip string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	f, err := os.OpenFile(s.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", ip)

	return nil
}

func (s *fileStorage) List() ([]string, error) {
	s.mutex.RLock()
	b, err := os.ReadFile(s.filePath)
	s.mutex.RUnlock()
	if err != nil {
		return nil, err
	}

	ips := strings.Split(string(b), "\n")
	if ips[len(ips)-1] == "" {
		ips = ips[:len(ips)-1]
	}

	return ips, nil
}

func (s *fileStorage) Close() error {
	return nil
}

func init() {
	storage.Register("file", openFileStorage)
}
