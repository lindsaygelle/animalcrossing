package species

import "testing"

func TestPLeo(t *testing.T) {
	var s string = "P. leo"
	if ok := pLeo == s; !ok {
		t.Fatalf("pLeo != %s", s)
	}
}
