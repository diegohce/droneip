package memstorage

import "testing"

func TestMemStorage(t *testing.T) {
	ms, _ := openMemStorage("3")

	ms.Save("1.1.1.1")
	ms.Save("2.2.2.2")
	ms.Save("3.3.3.3")
	ms.Save("4.4.4.4")

	l, _ := ms.List()

	if l[0] != "4.4.4.4" {
		t.Error("bad slice reordering")
	}

}
