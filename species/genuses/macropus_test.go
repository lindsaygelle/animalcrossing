package genuses

import "testing"

func TestMacropus(t *testing.T) {
	var s string = "Macropus"
	if ok := macropus == s; !ok {
		t.Fatalf("macropus != %s", s)
	}
}
