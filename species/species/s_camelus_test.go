package species

import "testing"

func TestSCamelus(t *testing.T) {
	var s string = "S. camelus"
	if ok := sCamelus == s; !ok {
		t.Fatalf("sCamelus != %s", s)
	}
}
