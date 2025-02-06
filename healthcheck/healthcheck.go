package healthcheck

import (
	"encoding/json"
	"net/http"
)

type Pinger interface {
	Ping() error
}

type HealthCheckFunc func() error

func (h HealthCheckFunc) Ping() error {
	return h()
}

func HealthCheck(pingers ...Pinger) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		response := struct {
			Status string   `json:"status"`
			Errors []string `json:"errors"`
		}{
			Status: "ok",
			Errors: []string{},
		}

		for _, p := range pingers {
			if err = p.Ping(); err != nil {
				response.Errors = append(response.Errors, err.Error())
			}
		}

		if len(response.Errors) > 0 {
			response.Status = "error"
		}

		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(&response); err != nil {
			w.Header().Del("Content-Type")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
