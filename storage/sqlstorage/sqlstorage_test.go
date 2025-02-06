package sqlstorage

import (
	"testing"
)

func TestSQLStorage(t *testing.T) {

	s1, _ := openSQLStorage("sqlite::memory:")
	defer s1.Close()

	s, err := openSQLStorage("sqlite::memory:")
	if err != nil {
		t.Error(err)
	}
	defer s.Close()

	s.Save("1.1.1.1")
	s.Save("1.1.1.1")
	s.List()
}
