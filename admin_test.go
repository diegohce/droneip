package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/diegohce/droneip/mxcache"
	"github.com/diegohce/droneip/storage"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
)

type MXRedisMock struct {
	cli  *redis.Client
	mock redismock.ClientMock
}

func (c *MXRedisMock) RedisClient() *redis.Client {
	return c.cli
}

func (c *MXRedisMock) Get(key string, i interface{}) error {
	result := c.cli.Get(context.TODO(), key)
	err := result.Err()
	if err != nil {
		if err != redis.Nil {
			return err
		}
		return mxcache.ErrNotFound
	}

	json.Unmarshal([]byte(result.Val()), i)

	return nil
}

func (c *MXRedisMock) Set(key string, data interface{}, ex int) error {
	return nil
}

func (c *MXRedisMock) Expire(pattern string) (mxcache.ExpiredKeys, error) {
	return mxcache.ExpiredKeys{}, nil
}

func (c *MXRedisMock) Incr(key string) (int64, error) {
	return 0, nil
}

func (c *MXRedisMock) IncrBy(key string, value int64) (int64, error) {
	return 0, nil
}

func (c *MXRedisMock) ExpireAt(key string, time time.Time) error {
	return nil
}

func (c *MXRedisMock) Ping() error {
	return nil
}

func TestAdminCentre(t *testing.T) {

	cli, mock := redismock.NewClientMock()

	mock.ExpectGet("droneip-10.0.0.1").SetVal(`{"ValidIP": false}`)
	mock.ExpectGet("droneip-10.0.0.2").SetVal(`{"validIP": true}`)
	mock.ExpectGet("droneip-10.0.0.3").SetErr(redis.Nil)

	red := MXRedisMock{
		cli:  cli,
		mock: mock,
	}
	store, _ := storage.Open("", "")
	ac := NewAdminCentre(&red, store)

	cases := []struct {
		IP   string
		want bool
	}{
		{"10.0.0.1", false},
		{"10.0.0.2", true},
		{"10.0.0.3", true},
	}

	for _, c := range cases {

		body := fmt.Sprintf(`{"addr": "%s"}`, c.IP)

		req := httptest.NewRequest(http.MethodPost, "/droneip/is_valid", strings.NewReader(body))
		res := httptest.NewRecorder()

		ac.ipIsValid(res, req)
		t.Log(res.Body.String())

		response := struct {
			IsValid bool `json:"is_valid"`
		}{}

		json.NewDecoder(res.Result().Body).Decode(&response)

		if response.IsValid != c.want {
			t.Errorf("%s: got %t want %t", c.IP, response.IsValid, c.want)
		}
	}

}

func TestAdminCentreHistory(t *testing.T) {

	cli, mock := redismock.NewClientMock()

	red := MXRedisMock{
		cli:  cli,
		mock: mock,
	}
	store, _ := storage.Open("", "")
	ac := NewAdminCentre(&red, store)

	req := httptest.NewRequest(http.MethodGet, "/droneip/history", nil)
	res := httptest.NewRecorder()

	ac.ipHistory(res, req)
	t.Log(res.Body.String())
}
