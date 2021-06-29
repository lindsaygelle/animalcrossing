package genuses

import "testing"

func TestGenusSus(t *testing.T) {
	var s string = "Sus"
	if ok := sus == s; !ok {
		t.Fatalf("sus != %s", s)
	}
}
