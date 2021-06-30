package orders

import "testing"

func TestOctopoda(t *testing.T) {
	var s string = "Octopoda"
	if ok := octopoda == s; !ok {
		t.Fatalf("octopoda != %s", s)
	}
}
