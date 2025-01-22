package version

import (
	"net/http/httptest"
	"testing"
)

func TestVersion(t *testing.T) {

	req := httptest.NewRequest("GET", "/version", nil)
	res := httptest.NewRecorder()

	Version(res, req)

}
