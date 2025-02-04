package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/diegohce/droneip/config"
)

func TestHappyTrail(t *testing.T) {

	config.FromEnv()
	config.FromEnvWithPrefix("TEST_")

	expected := []string{
		"localhost",
		"egorlami",
		"",
		"8080",
	}

	values := make([]string, 4)

	values[0] = config.Get("HOSTNAME")
	values[1] = config.Get("USERNAME")
	values[2] = config.Get("PASSWORD")
	values[3] = config.Get("PORT", "8080")

	for i, v := range values {
		if v != expected[i] {
			t.Error("expected", expected[i], "got", v)
		}
	}

	config.Set("PARAM", "pampam")

	if v := config.Get("PARAM"); v != "pampam" {
		t.Error("expected 'pampam' got", v)
	}

}

func TestNewValues(t *testing.T) {

	cfg := config.NewValues()

	cfg.FromEnv()
	cfg.FromEnvWithPrefix("TEST_")

	expected := []string{
		"localhost",
		"egorlami",
		"",
		"8080",
	}

	values := make([]string, 4)

	values[0] = cfg.Get("HOSTNAME")
	values[1] = cfg.Get("USERNAME")
	values[2] = cfg.Get("PASSWORD")
	values[3] = cfg.Get("PORT", "8080")

	for i, v := range values {
		if v != expected[i] {
			t.Error("expected", expected[i], "got", v)
		}
	}

	cfg.Set("PARAM", "pampam")

	if v := cfg.Get("PARAM"); v != "pampam" {
		t.Error("expected 'pampam' got", v)
	}

}

func TestGetInt(t *testing.T) {

	config.FromEnvWithPrefix("TEST_")

	if config.GetInt("USERID", 70) != 76 {
		t.Errorf("got %d want 76", config.GetInt("USERID", 70))
	}

}

func TestGetDuration(t *testing.T) {

	config.FromEnvWithPrefix("TEST_")

	if config.GetDuration("DURATION") != time.Duration(10*time.Second) {
		t.Errorf("got %d want 76", config.GetInt("USERID", 70))
	}

}

func TestMandatory(t *testing.T) {

	err := config.FromEnvWithPrefix("TEST_", "USERNAME", "MANDATORY")
	if err == nil {
		t.Errorf("got err == nil want mandatoy error")
	}

}

func TestMain(m *testing.M) {
	os.Setenv("TEST_HOSTNAME", "localhost")
	os.Setenv("TEST_USERNAME", "egorlami")
	os.Setenv("TEST_USERID", "76")
	os.Setenv("TEST_DURATION", "10s")
	os.Exit(m.Run())
}
