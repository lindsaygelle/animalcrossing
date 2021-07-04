package genuses

import "testing"

func TestUrsus(t *testing.T) {
	var s string = "Ursus"
	if ok := ursus == s; !ok {
		t.Fatalf("ursus != %s", s)
	}
}
