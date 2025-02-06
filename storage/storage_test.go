package storage

import (
	"testing"
)

func TestSQLStorageOK(t *testing.T) {

	s, err := Open("", "")
	if err != nil {
		t.Error(err)
	}
	defer s.Close()

	s.Save("1.1.1.1")
	s.Save("1.1.1.1")
	s.List()
}

func TestSQLStorageErr(t *testing.T) {

	_, err := Open("none", "")
	if err == nil {
		t.Error("got err want err == nil")
	}
}
