package genuses

import "testing"

func TestGenusGorilla(t *testing.T) {
	var s string = "Gorilla"
	if ok := gorilla == s; !ok {
		t.Fatalf("gorilla != %s", s)
	}
}
