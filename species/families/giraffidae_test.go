package families

import "testing"

func TestFamilyGiraffidae(t *testing.T) {
	var s string = "Giraffidae"
	if ok := giraffidae == s; !ok {
		t.Fatalf("giraffidae != %s", s)
	}
}
