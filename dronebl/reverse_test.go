package dronebl

import "testing"

func TestReverse(t *testing.T) {

	ip := "1.2.3.4"
	expected := "4.3.2.1"

	rev := reverseOctets(ip)

	if rev != expected {
		t.Errorf("got %s want %s", rev, expected)
	}

}
