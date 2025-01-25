package mxcache

import (
	"testing"
	"time"
)

func cacher(uri string) MXCacher {
	c, _ := NewMXCache(uri)
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
	}

	for _, c := range cases {
		c.cache.Set("key", "value", 0)
		c.cache.Get("key", "value")
		c.cache.ExpireAt("key", time.Now())
		c.cache.Incr("key")
		c.cache.IncrBy("key", 2)
		c.cache.Expire("key*")
		c.cache.Ping()
	}

}
