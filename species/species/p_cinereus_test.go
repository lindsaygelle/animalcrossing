package species

import "testing"

func TestPCinereus(t *testing.T) {
	var s string = "P. cinereus"
	if ok := pCinereus == s; !ok {
		t.Fatalf("pCinereus != %s", s)
	}
}
