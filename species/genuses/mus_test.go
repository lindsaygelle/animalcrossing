package genuses

import "testing"

func TestGenusMus(t *testing.T) {
	var s string = "Mus"
	if ok := mus == s; !ok {
		t.Fatalf("mus != %s", s)
	}
}
