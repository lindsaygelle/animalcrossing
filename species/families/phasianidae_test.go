package families

import "testing"

func TestFamilyPhasianidae(t *testing.T) {
	var s string = "Phasianidae"
	if ok := phasianidae == s; !ok {
		t.Fatalf("phasianidae != %s", s)
	}
}
