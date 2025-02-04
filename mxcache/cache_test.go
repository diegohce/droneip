package mxcache

import (
	"errors"
	"testing"
	"time"
)

func cacher(uri string) MXCacher {
	c, err := NewMXCache(uri)
	if err != nil {
		//c, _ = NewMXCache("")
	}
	return c
}

func TestCache(t *testing.T) {

	cases := []struct {
		cache MXCacher
	}{
		{cache: cacher("")},
		{cache: cacher("memory://")},
		{cache: cacher("redis://127.0.0.1:6379/1")},
		{cache: cacher("mem+redis://127.0.0.1:6379/1")},
		{cache: cacher("rediscluster://127.0.0.1:6379/1")},
		{cache: cacher("mem+rediscluster://127.0.0.1:6379/1")},
	}

	for _, c := range cases {

		if nilC, ok := c.cache.(nilCache); ok {

			nilC.Set("key", "value", 0)
			nilC.Get("key", "value")
			nilC.ExpireAt("key", time.Now())
			nilC.Incr("key")
			nilC.IncrBy("key", 2)
			nilC.Expire("key*")
			nilC.Ping()

		} else if memC, ok := c.cache.(*memoryCache); ok {

			memC.Set("key", "value", 0)
			memC.Get("key", "value")
			memC.ExpireAt("key", time.Now())
			memC.Incr("key")
			memC.IncrBy("key", 2)
			memC.Expire("key*")
			memC.Ping()

		} else if redisC, ok := c.cache.(*redisCache); ok {

			redisC.Set("key", "value", 0)
			redisC.Get("key", "value")
			redisC.ExpireAt("key", time.Now())
			redisC.Incr("key")
			redisC.IncrBy("key", 2)
			redisC.Expire("key*")
			redisC.Ping()

			redisC.RedisClient()
			redisC.Keys("")
			redisC.RemoveKeys("key*")

		} else if redisCC, ok := c.cache.(*redisClusterCache); ok {

			redisCC.Set("key", "value", 0)
			redisCC.Get("key", "value")
			redisCC.ExpireAt("key", time.Now())
			redisCC.Incr("key")
			redisCC.IncrBy("key", 2)
			redisCC.Expire("key*")
			redisCC.Ping()

			redisCC.RedisClient()
			redisCC.Keys("")
			redisCC.RemoveKeys("key*")

		} else if memRedisC, ok := c.cache.(*memRedisCache); ok {

			memRedisC.Set("key", "value", 0)
			memRedisC.Get("key", "value")
			memRedisC.ExpireAt("key", time.Now())
			memRedisC.Incr("key")
			memRedisC.IncrBy("key", 2)
			memRedisC.Expire("key*")
			memRedisC.Ping()

		} else if memRedisCC, ok := c.cache.(*memRedisClusterCache); ok {

			memRedisCC.Set("key", "value", 0)
			memRedisCC.Get("key", "value")
			memRedisCC.ExpireAt("key", time.Now())
			memRedisCC.Incr("key")
			memRedisCC.IncrBy("key", 2)
			memRedisCC.Expire("key*")
			memRedisCC.Ping()

		}

	}

}

func TestExpireErrors(t *testing.T) {
	var ee ExpireErrors

	ee = append(ee, errors.New("error1"), errors.New("error2"))

	s := ee.Error()

	if s != "error1, error2" {
		t.Error("bad error concat")
	}
}
