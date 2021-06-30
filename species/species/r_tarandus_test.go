package species

import "testing"

func TestRTarandus(t *testing.T) {
	var s string = "R. tarandus"
	if ok := rTarandus == s; !ok {
		t.Fatalf("rTarandus != %s", s)
	}
}
