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

type redisClusterCache struct {
	ctx   context.Context
	cache *redis.ClusterClient
	//options *redis.Options
}

func newRedisClusterCache(u *url.URL) (MXCacher, error) {

	switch u.Scheme {
	case "rediscluster":
		u.Scheme = "redis"
	case "redisclusters":
		u.Scheme = "rediss"
	}

	c := &redisClusterCache{}

	options, err := redis.ParseClusterURL(u.String())
	if err != nil {
		return nil, err
	}

	c.cache = redis.NewClusterClient(options)
	c.ctx = context.Background()

	err = c.cache.Ping(c.ctx).Err()

	return c, err
}

func (c *redisClusterCache) Ping() error {
	c.cache.Ping(c.ctx).Err()
	return nil
}

func (c *redisClusterCache) Get(key string, i interface{}) error {

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

func (c *redisClusterCache) Set(key string, data interface{}, ex int) error {
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

func (c *redisClusterCache) Expire(pattern string) (ExpiredKeys, error) {

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

func (c *redisClusterCache) Incr(key string) (int64, error) {
	result, err := c.cache.Incr(c.ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (c *redisClusterCache) IncrBy(key string, value int64) (int64, error) {
	result, err := c.cache.IncrBy(c.ctx, key, value).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (c *redisClusterCache) ExpireAt(key string, time time.Time) error {
	return c.cache.ExpireAt(c.ctx, key, time).Err()
}

func (c *redisClusterCache) RedisClient() *redis.ClusterClient {
	return c.cache
}

func (c *redisClusterCache) Keys(pattern string) ([]string, error) {
	if pattern == "" {
		pattern = "*"
	}

	return c.cache.Keys(c.ctx, pattern).Result()
}

func (c *redisClusterCache) RemoveKeys(keys ...string) error {
	return c.cache.Del(c.ctx, keys...).Err()
}
