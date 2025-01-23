//go:generate bash gen.sh

package version

import (
	"encoding/json"
	"net/http"
)

type versionHandler struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	When    string `json:"date"`
}

func (v *versionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(v)

}
