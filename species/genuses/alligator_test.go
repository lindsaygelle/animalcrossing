package genuses

import "testing"

func TestGenusAlligator(t *testing.T) {
	var s string = "Alligator"
	if ok := alligator == s; !ok {
		t.Fatalf("alligator != %s", s)
	}
}
