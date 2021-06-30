package genuses

import "testing"

func TestFelis(t *testing.T) {
	var s string = "Felis"
	if ok := felis == s; !ok {
		t.Fatalf("felis != %s", s)
	}
}
