package memstorage

import (
	"strconv"
	"sync"

	"github.com/diegohce/droneip/storage"
)

type memStorage struct {
	maxItems int
	items    []string
	mutex    sync.RWMutex
}

func openMemStorage(dsn string) (storage.Storager, error) {

	max, err := strconv.Atoi(dsn)
	if err != nil {
		return nil, err
	}

	fs := memStorage{
		maxItems: max,
		items:    make([]string, 0, max),
	}
	return &fs, nil
}

func (s *memStorage) Save(ip string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.items) == s.maxItems {
		for i := s.maxItems; i > 1; i-- {
			s.items[i-1] = s.items[i-2]
		}
		s.items[0] = ip
	} else {
		s.items = append(s.items, ip)
	}

	return nil
}

func (s *memStorage) List() ([]string, error) {
	ips := make([]string, s.maxItems)

	s.mutex.RLock()
	copy(ips, s.items)
	s.mutex.RUnlock()

	return ips, nil
}

func init() {
	storage.Register("mem", openMemStorage)
}
