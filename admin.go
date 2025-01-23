package main

import (
	"net/http"

	"github.com/diegohce/droneip/healthcheck"
	"github.com/diegohce/droneip/internal/version"

	"github.com/diegohce/droneip/ctcodecs"
	_ "github.com/diegohce/droneip/ctcodecs/allcodecs"

	"github.com/diegohce/droneip/logger"
	mx2 "github.com/diegohce/droneip/mxcache"

	"github.com/go-redis/redis/v8"
)

type redisCache interface {
	RedisClient() *redis.Client
}

type AdminCentre struct {
	cache mx2.MXCacher
	r     *http.ServeMux
}

func NewAdminCentre(cache mx2.MXCacher) *AdminCentre {

	r := http.ServeMux{}

	ac := AdminCentre{cache, &r}
	hc := healthcheck.HealthCheck(cache)

	ac.r.Handle("GET /droneip/keys", http.HandlerFunc(ac.cacheKeys))
	ac.r.Handle("GET /droneip/version", &version.Handler)
	ac.r.Handle("GET /droneip/health_check", hc)

	return &ac
}

func (a *AdminCentre) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.r.ServeHTTP(w, r)
}

func (a *AdminCentre) cacheKeys(w http.ResponseWriter, r *http.Request) {

	codec, err := ctcodecs.New(r.Header.Get("Content-Type"))
	if err != nil {
		codec, _ = ctcodecs.New("application/json")
	}

	redCache, ok := a.cache.(redisCache)
	if !ok {
		logger.LogError("cache keys only implemented for redis cache").Write()
		http.Error(w, "cache keys only implemented for redis cache", http.StatusBadRequest)
		return
	}

	rc := redCache.RedisClient()

	ctx := r.Context()
	result := rc.Keys(ctx, "droneip-*")
	if result.Err() != nil {
		logger.LogError("error getting keys from cache backend", "error", result.Err().Error()).Write()
		http.Error(w, "error getting keys from cache backend: "+result.Err().Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Keys []string `json:"keys"`
	}{
		Keys: result.Val(),
	}

	codec.NewEncoder(w).Encode(response)
}
