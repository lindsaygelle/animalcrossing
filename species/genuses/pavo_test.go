package genuses

import "testing"

func TestGenusPavo(t *testing.T) {
	var s string = "Pavo"
	if ok := pavo == s; !ok {
		t.Fatalf("pavo != %s", s)
	}
}
