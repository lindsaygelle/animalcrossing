package families

import "testing"

func TestFamilyFelidae(t *testing.T) {
	var s string = "Felidae"
	if ok := felidae == s; !ok {
		t.Fatalf("felidae != %s", s)
	}
}
