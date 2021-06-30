package orders

import "testing"

func TestSquamata(t *testing.T) {
	var s string = "Squamata"
	if ok := squamata == s; !ok {
		t.Fatalf("squamata != %s", s)
	}
}
