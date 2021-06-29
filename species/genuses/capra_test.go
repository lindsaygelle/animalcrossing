package genuses

import "testing"

func TestGenusCapra(t *testing.T) {
	var s string = "Capra"
	if ok := capra == s; !ok {
		t.Fatalf("capra != %s", s)
	}
}
