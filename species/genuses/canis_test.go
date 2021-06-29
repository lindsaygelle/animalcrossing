package genuses

import "testing"

func TestGenusCanis(t *testing.T) {
	var s string = "Canis"
	if ok := canis == s; !ok {
		t.Fatalf("canis != %s", s)
	}
}
