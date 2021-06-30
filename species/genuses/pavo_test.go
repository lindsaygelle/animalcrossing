package genuses

import "testing"

func TestPavo(t *testing.T) {
	var s string = "Pavo"
	if ok := pavo == s; !ok {
		t.Fatalf("pavo != %s", s)
	}
}
