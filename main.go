//go:build !test

package main

import (
	"net/http"
	"os"
	"time"

	"github.com/diegohce/droneip/config"
	"github.com/diegohce/droneip/logger"

	"github.com/diegohce/droneip/storage"
	_ "github.com/diegohce/droneip/storage/filestorage"
	_ "github.com/diegohce/droneip/storage/httpstorage"
	_ "github.com/diegohce/droneip/storage/memstorage"
	_ "github.com/diegohce/droneip/storage/sqlstorage"

	mx2 "github.com/diegohce/droneip/mxcache"
)

func main() {

	err := config.FromEnvWithPrefix("DRONEIP_", "DESTINATION_URL")
	if err != nil {
		logger.LogError("config error", "err", err.Error()).Write()
		os.Exit(1)
	}

	store, err := storage.Open(config.Get("STORAGE_TYPE"), config.Get("STORAGE_CONFIG"))
	if err != nil {
		logger.LogError("error starting storage", "err", err.Error()).Write()
		os.Exit(1)
	}

	cache, err := mx2.NewMXCache(config.Get("CACHE_URL", ""))
	if err != nil {
		logger.LogError("error starting cache", "err", err.Error()).Write()
		os.Exit(1)
	}

	ac := NewAdminCentre(cache, store)
	go http.ListenAndServe(config.Get("ADMIN_BIND", ":8081"), ac)

	ttl, err := time.ParseDuration(config.Get("CACHE_TTL", "24h"))
	if err != nil {
		ttl, _ = time.ParseDuration("24h")
	}
	handler := DroneHandler{
		cache:    cache,
		cacheTTL: int(ttl.Seconds()),
		store:    store,
	}

	http.Handle("/", &handler)

	logger.LogInfo("starting droneip MiM",
		"bind", config.Get("BIND", ":8080"),
		"forwarding_to", config.Get("DESTINATION_URL")).Write()

	http.ListenAndServe(config.Get("BIND", ":8080"), nil)
}
