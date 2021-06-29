package classes

import "testing"

func TestSauropsida(t *testing.T) {
	var s string = "Sauropsida"
	if ok := sauropsida == s; !ok {
		t.Fatalf("sauropsida != %s", s)
	}
}
