package families

import "testing"

func TestFelidae(t *testing.T) {
	var s string = "Felidae"
	if ok := felidae == s; !ok {
		t.Fatalf("felidae != %s", s)
	}
}
