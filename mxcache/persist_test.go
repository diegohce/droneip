package mxcache_test

import (
	"testing"

	"github.com/diegohce/droneip/mxcache"
)

func Test00PersistenceSet(t *testing.T) {

	cache, err := mxcache.NewMXCache("memory://mem/?persist=cache.dat")
	if err != nil {
		t.Fatal(err)
	}

	cache.Set("name", "Diego", 3600)
}

func Test01PersistenceGet(t *testing.T) {

	cache, err := mxcache.NewMXCache("memory://mem/?persist=cache.dat")
	if err != nil {
		t.Fatal(err)
	}

	name := ""

	cache.Get("name", name)

	t.Log(name)
}
