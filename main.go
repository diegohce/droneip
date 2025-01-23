//go:build !test

package main

import (
	"net/http"
	"os"
	"time"

	"github.com/diegohce/droneip/config"
	"github.com/diegohce/droneip/logger"

	mx2 "github.com/diegohce/droneip/mxcache"
)

func main() {

	config.FromEnvWithPrefix("DRONEIP_")

	if config.Get("DESTINATION_URL", "NN") == "NN" {
		logger.LogError("destination URL not set").Write()
		os.Exit(1)
	}

	cache, err := mx2.NewMXCache(config.Get("CACHE_URL", ""))
	if err != nil {
		logger.LogError("error starting cache", "err", err.Error()).Write()
		os.Exit(1)
	}

	ac := NewAdminCentre(cache)
	go http.ListenAndServe(config.Get("ADMIN_BIND", ":8081"), ac)

	ttl, err := time.ParseDuration(config.Get("CACHE_TTL", "24h"))
	if err != nil {
		ttl, _ = time.ParseDuration("24h")
	}
	handler := DroneHandler{
		cache:    cache,
		cacheTTL: int(ttl.Seconds()),
	}

	http.Handle("/", &handler)

	logger.LogInfo("starting droneip MiM",
		"bind", config.Get("BIND", ":8080"),
		"forwarding_to", config.Get("DESTINATION_URL")).Write()

	http.ListenAndServe(config.Get("BIND", ":8080"), nil)
}
