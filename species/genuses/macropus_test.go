package genuses

import "testing"

func TestGenusMacropus(t *testing.T) {
	var s string = "Macropus"
	if ok := macropus == s; !ok {
		t.Fatalf("macropus != %s", s)
	}
}
