package mxcache

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"
)

var ErrNotFound = errors.New("not found")

type MXCacheCreator func(u *url.URL) (MXCacher, error)

type ExpiredKeys []string

type MXCacher interface {
	Get(key string, i interface{}) error
	Set(key string, data interface{}, ex int) error
	Expire(pattern string) (ExpiredKeys, error)
	Incr(key string) (int64, error)
	IncrBy(key string, value int64) (int64, error)
	ExpireAt(key string, time time.Time) error
	Ping() error
}

type MXKeysLister interface {
	Keys(pattern string) ([]string, error)
}

type MXKeysRemover interface {
	RemoveKeys(keys ...string) error
}

var cacheBackends = map[string]MXCacheCreator{
	"memory":            newMemoryCache,
	"redis":             newRedisCache,
	"rediss":            newRedisCache,
	"mem+redis":         newMemRedisCache,
	"mem+rediss":        newMemRedisCache,
	"rediscluster":      newRedisClusterCache,
	"redisclusters":     newRedisClusterCache,
	"mem+rediscluster":  newMemRedisClusterCache,
	"mem+redisclusters": newMemRedisClusterCache,
}

func NewMXCache(uri string) (MXCacher, error) {
	if uri == "" {
		return nilCache{}, nil
	}
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	backendCreator, ok := cacheBackends[u.Scheme]
	if !ok {
		return nil, fmt.Errorf("invalid cache backend %s", u.Scheme)
	}
	log.Println("Setting up", u.Scheme, "cache with", u.String())
	return backendCreator(u)
}

type nilCache struct{}

func (c nilCache) Ping() error {
	return nil
}

func (c nilCache) Set(key string, data interface{}, ex int) error {
	return nil
}

func (c nilCache) Get(key string, i interface{}) error {
	return nil
}

func (c nilCache) Expire(pattern string) (ExpiredKeys, error) {
	return []string{}, nil
}

func (c nilCache) Incr(key string) (int64, error) {
	return 0, nil
}

func (c nilCache) IncrBy(key string, value int64) (int64, error) {
	return 0, nil
}

func (c nilCache) ExpireAt(key string, time time.Time) error {
	return nil
}
