package classes

import "testing"

func TestAves(t *testing.T) {
	var s string = "Aves"
	if ok := aves == s; !ok {
		t.Fatalf("aves != %s", s)
	}
}
