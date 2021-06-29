package families

import "testing"

func TestPhasianidae(t *testing.T) {
	var s string = "Phasianidae"
	if ok := phasianidae == s; !ok {
		t.Fatalf("phasianidae != %s", s)
	}
}
