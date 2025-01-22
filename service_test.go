package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/diegohce/config"
	mx2 "github.com/diegohce/mxcache"
)

func TestService(t *testing.T) {

	remoteServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	defer remoteServer.Close()

	config.Values.Set("INSPECT_HEADER", "X-TestHeader")
	config.Values.Set("DESTINATION_URL", remoteServer.URL)

	cache, _ := mx2.NewMXCache("memory://")

	cases := []struct {
		method         string
		ip             string
		responseStatus int
	}{
		{"POST", "35.172.175.100", http.StatusOK},
		{"POST", "193.56.64.251", http.StatusTeapot},

		//secod pass to try cache
		{"POST", "35.172.175.100", http.StatusOK},
		{"POST", "193.56.64.251", http.StatusTeapot},
	}

	h := &DroneHandler{
		cache:    cache,
		cacheTTL: 24 * 60 * 60,
	}

	for i, c := range cases {

		req := httptest.NewRequest(c.method, "/", strings.NewReader("hola mundo"))
		req.Header.Set("X-TestHeader", c.ip)

		res := httptest.NewRecorder()

		h.ServeHTTP(res, req)

		if res.Result().StatusCode != c.responseStatus {
			t.Fatalf("case %d: got status %d want status %d", i, res.Result().StatusCode, c.responseStatus)
		}

	}

}

func TestAdmin(t *testing.T) {

	cases := []struct {
		cacheURI string
		status   int
	}{
		{"memory://", 400},
		{"ER:redis://:@127.0.0.1:6378/0", 400},
		{"OK:redis://:@127.0.0.1:6379/0", 200},
	}

	for i, c := range cases {
		var cache mx2.MXCacher

		if strings.HasPrefix(c.cacheURI, "memory") {
			cache, _ = mx2.NewMXCache(c.cacheURI)

		} else {
			cache, _ = newFakeRedis(c.cacheURI)

		}

		admin := NewAdminCentre(cache)

		req := httptest.NewRequest("GET", "/droneip/keys", nil)
		res := httptest.NewRecorder()

		admin.ServeHTTP(res, req)

		if res.Result().StatusCode != c.status {
			t.Fatalf("case %d: got status %d want status %d", i, res.Result().StatusCode, c.status)
		}

	}

}
