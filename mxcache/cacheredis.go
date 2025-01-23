package mxcache

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	ctx   context.Context
	cache *redis.Client
	//options *redis.Options
}

func newRedisCache(u *url.URL) (MXCacher, error) {

	// redis://:master@127.0.0.1/1

	c := &redisCache{}

	options, err := redis.ParseURL(u.String())
	if err != nil {
		return nil, err
	}

	c.cache = redis.NewClient(options)
	c.ctx = context.Background()

	err = c.cache.Ping(c.ctx).Err()

	return c, err
}

func (c *redisCache) Ping() error {
	c.cache.Ping(c.ctx).Err()
	return nil
}

func (c *redisCache) Get(key string, i interface{}) error {

	result := c.cache.Get(c.ctx, key)
	err := result.Err()
	if err != nil {
		if err != redis.Nil {
			return err
		}
		return ErrNotFound
	}

	val, err := result.Bytes()
	if err != nil {
		return err
	}

	if err := json.NewDecoder(bytes.NewReader(val)).Decode(i); err != nil {
		return err
	}

	return nil
}

func (c *redisCache) Set(key string, data interface{}, ex int) error {
	ttl := time.Duration(0)
	if ex > 0 {
		ttl = time.Duration(ex) * time.Second
	}

	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return err
	}

	return c.cache.Set(c.ctx, key, buf.Bytes(), ttl).Err()
}

func (c *redisCache) Expire(pattern string) (ExpiredKeys, error) {

	var exkeys ExpiredKeys

	if !strings.Contains(pattern, "*") {
		exkeys = append(exkeys, pattern)
		return exkeys, c.cache.Del(c.ctx, pattern).Err()
	}

	keys, err := c.cache.Keys(c.ctx, pattern).Result()
	if err != nil {
		return exkeys, err
	}
	if len(keys) == 0 {
		return exkeys, nil
	}

	return keys, c.cache.Del(c.ctx, keys...).Err()
}

func (c *redisCache) Incr(key string) (int64, error) {
	result, err := c.cache.Incr(c.ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (c *redisCache) IncrBy(key string, value int64) (int64, error) {
	result, err := c.cache.IncrBy(c.ctx, key, value).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (c *redisCache) ExpireAt(key string, time time.Time) error {
	return c.cache.ExpireAt(c.ctx, key, time).Err()
}

func (c *redisCache) RedisClient() *redis.Client {
	return c.cache
}

func (c *redisCache) Keys(pattern string) ([]string, error) {
	if pattern == "" {
		pattern = "*"
	}

	return c.cache.Keys(c.ctx, pattern).Result()
}
func (c *redisCache) RemoveKeys(keys ...string) error {
	return c.cache.Del(c.ctx, keys...).Err()
}
