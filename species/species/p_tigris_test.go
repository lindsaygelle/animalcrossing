package species

import "testing"

func TestPTigris(t *testing.T) {
	var s string = "P. tigris"
	if ok := pTigris == s; !ok {
		t.Fatalf("pTigris != %s", s)
	}
}
