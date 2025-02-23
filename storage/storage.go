package storage

import "errors"

type Storager interface {
	Save(ip string) error
	List() ([]string, error)
	Close() error
}

type OpenStorageFunc func(string) (Storager, error)

var (
	operators = map[string]OpenStorageFunc{}
)

func Register(name string, fn OpenStorageFunc) {
	operators[name] = fn
}

func Open(name string, dsn string) (Storager, error) {
	if name == "" {
		return &nilStorage{}, nil
	}

	fn, exists := operators[name]
	if !exists {
		return nil, errors.New("invalid storage name")
	}
	return fn(dsn)
}

type nilStorage struct{}

func (s *nilStorage) Save(ip string) error {
	return nil
}
func (s *nilStorage) List() ([]string, error) {
	return []string{}, nil
}

func (s *nilStorage) Close() error {
	return nil
}
