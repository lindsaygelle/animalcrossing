package genuses

import "testing"

func TestGenusPhascolarctos(t *testing.T) {
	var s string = "Phascolarctos"
	if ok := phascolarctos == s; !ok {
		t.Fatalf("phascolarctos != %s", s)
	}
}
