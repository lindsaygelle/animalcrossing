package classes

import "testing"

func TestAmmalia(t *testing.T) {
	var s string = "Ammalia"
	if ok := ammalia == s; !ok {
		t.Fatalf("ammalia != %s", s)
	}
}
