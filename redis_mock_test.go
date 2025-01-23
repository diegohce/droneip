package main

import (
	"errors"
	"time"

	mx2 "github.com/diegohce/droneip/mxcache"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

type FakeRedis struct {
	cli  *redis.Client
	mock redismock.ClientMock
}

func newFakeRedis(uri string) (mx2.MXCacher, error) {
	cli, mock := redismock.NewClientMock()

	c := FakeRedis{
		cli:  cli,
		mock: mock,
	}

	if uri[:2] == "ER" {
		c.mock.ExpectKeys("droneip-*").SetErr(errors.New("error connecting to redis"))
	} else {
		c.mock.ExpectKeys("droneip-*").SetVal([]string{})
	}

	return &c, nil

}

func (c *FakeRedis) RedisClient() *redis.Client {
	return c.cli
}

func (c *FakeRedis) Get(key string, i interface{}) error {
	return nil
}

func (c *FakeRedis) Set(key string, data interface{}, ex int) error {
	return nil
}

func (c *FakeRedis) Expire(pattern string) (mx2.ExpiredKeys, error) {
	return mx2.ExpiredKeys{}, nil
}

func (c *FakeRedis) Incr(key string) (int64, error) {
	return 0, nil
}

func (c *FakeRedis) IncrBy(key string, value int64) (int64, error) {
	return 0, nil
}

func (c *FakeRedis) ExpireAt(key string, time time.Time) error {
	return nil
}

func (c *FakeRedis) Ping() error {
	return nil
}
