package httpstorage

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpStorage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := struct {
			IP string `json:"ip"`
		}{}

		json.NewDecoder(r.Body).Decode(&req)

		if req.IP != "1.1.1.1" {
			t.Errorf("got %s want 1.1.1.1", req.IP)
		}

	}))
	defer server.Close()

	hs, _ := openHttpStorage(server.URL + "/")
	defer hs.Close()

	err := hs.Save("1.1.1.1")
	if err != nil {
		t.Error(err)
	}

	hs.List()
}
