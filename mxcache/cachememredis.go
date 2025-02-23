package mxcache

import (
	"log"
	"net/url"
	"strings"
	"time"
)

type memRedisCache struct {
	mem   MXCacher
	redis MXCacher
}

func newMemRedisCache(u *url.URL) (MXCacher, error) {

	// mem+redis://:master@127.0.0.1/1

	c := &memRedisCache{}

	c.mem, _ = newMemoryCache(u)

	redisUri, _ := url.Parse(strings.SplitN(u.String(), "+", 2)[1])

	if r, err := newRedisCache(redisUri); err != nil {
		return nil, err
	} else {
		c.redis = r
	}

	return c, nil
}

func (c *memRedisCache) Ping() error {
	var err error

	if err = c.mem.Ping(); err != nil {
		return err
	}
	return c.redis.Ping()
}

func (c *memRedisCache) Get(key string, i interface{}) error {

	c.mem.Get(key, i)
	if i != nil {
		log.Println("mem+redis: mem")
		return nil
	}

	err := c.redis.Get(key, i)
	if err != nil {
		return err
	}
	if i == nil {
		return ErrNotFound
	}

	rc := c.redis.(*redisCache)
	keyTTL := rc.RedisClient().TTL(rc.ctx, key).Val()

	c.mem.Set(key, i, int(keyTTL))

	log.Println("mem+redis: redis. TTL:", keyTTL)
	return nil
}

func (c *memRedisCache) Set(key string, data interface{}, ex int) error {
	c.mem.Set(key, data, ex)
	return c.redis.Set(key, data, ex)
}

func (c *memRedisCache) Expire(pattern string) (ExpiredKeys, error) {
	c.mem.Expire(pattern)
	return c.redis.Expire(pattern)
}

func (c *memRedisCache) Incr(key string) (int64, error) {
	return 0, nil
}

func (c *memRedisCache) IncrBy(key string, value int64) (int64, error) {
	return 0, nil
}

func (c *memRedisCache) ExpireAt(key string, time time.Time) error {
	return nil
}
