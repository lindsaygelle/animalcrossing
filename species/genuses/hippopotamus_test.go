package genuses

import "testing"

func TestHippopotamus(t *testing.T) {
	var s string = "Hippopotamus"
	if ok := hippopotamus == s; !ok {
		t.Fatalf("hippopotamus != %s", s)
	}
}
