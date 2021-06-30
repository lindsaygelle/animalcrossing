package genuses

import "testing"

func TestBos(t *testing.T) {
	var s string = "Bos"
	if ok := bos == s; !ok {
		t.Fatalf("bos != %s", s)
	}
}
