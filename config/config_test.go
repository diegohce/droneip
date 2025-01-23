package config_test

import (
	"os"
	"testing"

	"github.com/diegohce/droneip/config"
)

func TestHappyTrail(t *testing.T) {

	config.Values.FromEnv()
	config.Values.FromEnvWithPrefix("TEST_")

	expected := []string{
		"localhost",
		"egorlami",
		"",
		"8080",
	}

	values := make([]string, 4)

	values[0] = config.Values.Get("HOSTNAME")
	values[1] = config.Values.Get("USERNAME")
	values[2] = config.Values.Get("PASSWORD")
	values[3] = config.Values.Get("PORT", "8080")

	for i, v := range values {
		if v != expected[i] {
			t.Fatal("expected", expected[i], "got", v)
		}
	}

	config.Values.Set("PARAM", "pampam")

	if v := config.Values.Get("PARAM"); v != "pampam" {
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
