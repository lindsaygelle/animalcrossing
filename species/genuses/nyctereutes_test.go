package genuses

import "testing"

func TestGenusNyctereutes(t *testing.T) {
	var s string = "Nyctereutes"
	if ok := nyctereutes == s; !ok {
		t.Fatalf("nyctereutes != %s", s)
	}
}
