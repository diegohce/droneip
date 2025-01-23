package config_test

import (
	"os"
	"testing"

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
			t.Fatal("expected", expected[i], "got", v)
		}
	}

	config.Set("PARAM", "pampam")

	if v := config.Get("PARAM"); v != "pampam" {
		t.Fatal("expected 'pampam' got", v)
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
			t.Fatal("expected", expected[i], "got", v)
		}
	}

	cfg.Set("PARAM", "pampam")

	if v := cfg.Get("PARAM"); v != "pampam" {
		t.Fatal("expected 'pampam' got", v)
	}

}

func TestMain(m *testing.M) {
	os.Setenv("TEST_HOSTNAME", "localhost")
	os.Setenv("TEST_USERNAME", "egorlami")
	os.Exit(m.Run())
}
