package domains

import "testing"

func TestEukarya(t *testing.T) {
	var s string = "Eukarya"
	if ok := eukarya == s; !ok {
		t.Fatalf("eukarya != %s", s)
	}
}
