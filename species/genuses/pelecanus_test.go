package genuses

import "testing"

func TestPelecanus(t *testing.T) {
	var s string = "Pelecanus"
	if ok := pelecanus == s; !ok {
		t.Fatalf("pelecanus != %s", s)
	}
}
