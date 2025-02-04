package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/diegohce/droneip/dronebl"
	"github.com/diegohce/droneip/storage"

	"github.com/diegohce/droneip/config"
	mx2 "github.com/diegohce/droneip/mxcache"

	"github.com/diegohce/droneip/logger"
)

type CacheValue struct {
	ValidIP bool
}

type DroneHandler struct {
	cache    mx2.MXCacher
	cacheTTL int
	store    storage.Storager
}

func (h *DroneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	remoteIP := strings.SplitN(r.RemoteAddr, ":", 2)[0]

	inspectHeader := config.Get("INSPECT_HEADER")
	if inspectHeader != "" {
		remoteIP = getRemoteIP(r.Header.Get(inspectHeader))
	}
	cacheKey := fmt.Sprintf("droneip-%s", remoteIP)

	var cv CacheValue

	err := h.cache.Get(cacheKey, &cv)
	if errors.Is(err, mx2.ErrNotFound) {

		if err := dronebl.Probe(remoteIP); err != nil {
			logger.LogInfo("ip exists in dronebl DB",
				"ip", remoteIP, "source", "dronebl.org").Write()

			cv.ValidIP = false
			h.cache.Set(cacheKey, &cv, h.cacheTTL)
			w.WriteHeader(http.StatusTeapot)
			return
		}
		cv.ValidIP = true
		h.cache.Set(cacheKey, &cv, h.cacheTTL)
	}

	if !cv.ValidIP {
		logger.LogInfo("ip exists in dronebl DB",
			"ip", remoteIP, "source", "cache").Write()

		h.store.Save(remoteIP)
		w.WriteHeader(http.StatusTeapot)
		return
	}

	destURL := config.Get("DESTINATION_URL")

	req, err := http.NewRequest(r.Method, destURL, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	req.Header = r.Header

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	io.Copy(w, res.Body)

}

func getRemoteIP(ips string) string {
	if !strings.Contains(ips, ",") {
		return ips
	}

	ipList := strings.Split(ips, ",")

	return strings.Trim(ipList[0], " ")
}
