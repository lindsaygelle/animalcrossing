package species

import "testing"

func TestCAegagrus(t *testing.T) {
	var s string = "C. aegagrus"
	if ok := cAegagrus == s; !ok {
		t.Fatalf("cAegagrus != %s", s)
	}
}
