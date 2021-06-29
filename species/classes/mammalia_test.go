package classes

import "testing"

func TestMammalia(t *testing.T) {
	var s string = "Mammalia"
	if ok := mammalia == s; !ok {
		t.Fatalf("mammalia != %s", s)
	}
}
