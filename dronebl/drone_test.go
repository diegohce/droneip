package dronebl

import "testing"

func TestDrone(t *testing.T) {

	cases := []struct {
		ip    string
		found bool
	}{
		{"193.56.64.251", true},
		{"35.172.175.100", false},
	}

	for _, c := range cases {
		err := Probe(c.ip)
		if (err != nil) != c.found {
			t.Fatalf("got ip %s found %t want found %t", c.ip, (err != nil), c.found)
		}
	}

}
