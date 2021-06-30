package orders

import "testing"

func TestAnura(t *testing.T) {
	var s string = "Anura"
	if ok := anura == s; !ok {
		t.Fatalf("anura != %s", s)
	}
}
