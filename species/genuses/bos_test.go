package genuses

import "testing"

func TestGenusBos(t *testing.T) {
	var s string = "Bos"
	if ok := bos == s; !ok {
		t.Fatalf("bos != %s", s)
	}
}
