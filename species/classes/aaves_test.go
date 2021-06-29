package classes

import "testing"

func TestAaves(t *testing.T) {
	var s string = "Aaves"
	if ok := aaves == s; !ok {
		t.Fatalf("aaves != %s", s)
	}
}
