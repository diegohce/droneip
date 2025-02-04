package filestorage

import (
	"os"
	"testing"
)

func TestFileStorage(t *testing.T) {
	fs, _ := openFileStorage("file:///tmp/droneip-storage.txt")
	defer os.Remove("/tmp/droneip-storage.txt")

	fs.Save("1.1.1.1")
	fs.Save("2.2.2.2")
	fs.Save("3.3.3.3")
	fs.Save("4.4.4.4")

	l, _ := fs.List()

	if len(l) != 4 {
		t.Error("bad file storge")
	}

}
