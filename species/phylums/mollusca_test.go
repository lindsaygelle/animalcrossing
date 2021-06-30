package phylums

import "testing"

func TestPhylumMollusca(t *testing.T) {
	var s string = "Mollusca"
	if ok := mollusca == s; !ok {
		t.Fatalf("mollusca != %s", s)
	}
}
