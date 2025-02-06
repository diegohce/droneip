package healthcheck

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
)

type mockDB struct{}

func (db *mockDB) Ping() error {
	return nil
}

type mockCache struct{}

func (db *mockCache) Ping() error {
	return nil
}

type mockSomeService struct{}

func (s *mockSomeService) Ping() error {
	return nil
}

func TestHealth(t *testing.T) {

	handler := HealthCheck(&mockDB{}, &mockCache{}, &mockSomeService{})

	rq := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, rq)

}

func TestHealthWithoutDependencies(t *testing.T) {

	handler := HealthCheck()

	rq := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, rq)

}

type errService struct{}

func (s *errService) Ping() error {
	return errors.New("bad service")
}

func TestHealthNotOK(t *testing.T) {

	handler := HealthCheck(&errService{}, &mockDB{}, &mockCache{}, &mockSomeService{})

	rq := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, rq)

	response := struct {
		Status string   `json:"status"`
		Errors []string `json:"errors"`
	}{}

	json.NewDecoder(rr.Result().Body).Decode(&response)

	if response.Status != "error" {
		t.Errorf("status: got %s want error", response.Status)
	}

}

func TestHealthCheckFunc(t *testing.T) {

	hc := HealthCheckFunc(func() error {
		return nil
	})

	err := hc.Ping()
	if err != nil {
		t.Error(err)
	}

}
