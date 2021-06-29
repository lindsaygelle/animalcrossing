package classes

import "testing"

func TestReptilia(t *testing.T) {
	var s string = "Reptilia"
	if ok := reptilia == s; !ok {
		t.Fatalf("reptilia != %s", s)
	}
}
