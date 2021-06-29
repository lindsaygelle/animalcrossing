package classes

import "testing"

func TestCephalopoda(t *testing.T) {
	var s string = "Cephalopoda"
	if ok := cephalopoda == s; !ok {
		t.Fatalf("cephalopoda != %s", s)
	}
}
