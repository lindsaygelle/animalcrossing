package genuses

import "testing"

func TestNyctereutes(t *testing.T) {
	var s string = "Nyctereutes"
	if ok := nyctereutes == s; !ok {
		t.Fatalf("nyctereutes != %s", s)
	}
}
