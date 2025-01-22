//go:generate bash gen.sh

package version

import (
	"encoding/json"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {

	ver := struct {
		Version string `json:"version"`
		Commit  string `json:"commit"`
		When    string `json:"date"`
	}{
		Version: VERSION,
		Commit:  COMMIT,
		When:    WHEN,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(ver)

}
