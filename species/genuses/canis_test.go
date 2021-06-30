package genuses

import "testing"

func TestCanis(t *testing.T) {
	var s string = "Canis"
	if ok := canis == s; !ok {
		t.Fatalf("canis != %s", s)
	}
}
