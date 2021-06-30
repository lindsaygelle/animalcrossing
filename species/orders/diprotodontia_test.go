package orders

import "testing"

func TestDiprotodontia(t *testing.T) {
	var s string = "Diprotodontia"
	if ok := diprotodontia == s; !ok {
		t.Fatalf("diprotodontia != %s", s)
	}
}
