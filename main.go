//go:build !test

package main

import (
	"net/http"
	"time"

	"github.com/diegohce/config"
	"github.com/diegohce/logger"

	mx2 "github.com/diegohce/mxcache"
)

func main() {

	config.Values.FromEnvWithPrefix("DRONEIP_")

	if config.Values.Get("DESTINATION_URL", "NN") == "NN" {
		logger.LogError("destination URL not set").Write()
		return
	}

	cache, _ := mx2.NewMXCache(config.Values.Get("CACHE_URL", ""))

	ac := NewAdminCentre(cache)
	go http.ListenAndServe(config.Values.Get("ADMIN_BIND", ":8081"), ac)

	ttl, err := time.ParseDuration(config.Values.Get("CACHE_TTL", "24h"))
	if err != nil {
		ttl, _ = time.ParseDuration("24h")
	}
	handler := DroneHandler{
		cache:    cache,
		cacheTTL: int(ttl.Seconds()),
	}

	http.Handle("/", &handler)

	logger.LogInfo("starting droneip MiM",
		"bind", config.Values.Get("BIND", ":8080"),
		"forwarding_to", config.Values.Get("DESTINATION_URL")).Write()

	http.ListenAndServe(config.Values.Get("BIND", ":8080"), nil)
}
