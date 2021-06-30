package genuses

import "testing"

func TestSus(t *testing.T) {
	var s string = "Sus"
	if ok := sus == s; !ok {
		t.Fatalf("sus != %s", s)
	}
}
