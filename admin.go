package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/diegohce/droneip/config"
	"github.com/diegohce/droneip/dronebl"
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
	ac.r.Handle("POST /droneip/is_valid", http.HandlerFunc(ac.ipIsValid))

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

func (a *AdminCentre) ipIsValid(w http.ResponseWriter, r *http.Request) {

	rq := struct {
		Addr string `json:"addr"`
	}{}

	ttl, err := time.ParseDuration(config.Get("CACHE_TTL", "24h"))
	if err != nil {
		ttl, _ = time.ParseDuration("24h")
	}

	codec, err := ctcodecs.New(r.Header.Get("Content-Type"))
	if err != nil {
		codec, _ = ctcodecs.New("application/json")
	}

	if err = codec.NewDecoder(r.Body).Decode(&rq); err != nil {
		logger.LogError("error decoding request body", "err", err.Error()).Write()
		http.Error(w, "error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	cacheKey := fmt.Sprintf("droneip-%s", rq.Addr)

	cv := CacheValue{}

	err = a.cache.Get(cacheKey, &cv)
	if errors.Is(err, mx2.ErrNotFound) {
		if err := dronebl.Probe(rq.Addr); err != nil {
			logger.LogInfo("ip exists in dronebl DB",
				"ip", rq.Addr, "source", "dronebl.org").Write()

			cv.ValidIP = false
		} else {
			cv.ValidIP = true
		}

		a.cache.Set(cacheKey, &cv, int(ttl.Seconds()))
	}

	res := struct {
		IsValid bool `json:"is_valid"`
	}{
		IsValid: cv.ValidIP,
	}

	codec.NewEncoder(w).Encode(&res)
}
