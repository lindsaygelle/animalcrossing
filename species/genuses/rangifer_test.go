package genuses

import "testing"

func TestGenusRangifer(t *testing.T) {
	var s string = "Rangifer"
	if ok := rangifer == s; !ok {
		t.Fatalf("rangifer != %s", s)
	}
}
