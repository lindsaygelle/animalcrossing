package genuses

import "testing"

func TestEquus(t *testing.T) {
	var s string = "Equus"
	if ok := equus == s; !ok {
		t.Fatalf("equus != %s", s)
	}
}
