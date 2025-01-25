package mxcache

import (
	"testing"
)

func Test00PersistenceSet(t *testing.T) {

	cache, err := NewMXCache("memory://mem/?persist=cache.dat")
	if err != nil {
		t.Error(err)
	}

	cache.Set("name", "Diego", 3600)
}

func Test01PersistenceGet(t *testing.T) {

	cache, err := NewMXCache("memory://mem/?persist=cache.dat")
	if err != nil {
		t.Error(err)
	}

	name := ""

	cache.Get("name", name)

	t.Log(name)
}
