package families

import "testing"

func TestFamilyAlligatoridae(t *testing.T) {
	var s string = "Alligatoridae"
	if ok := alligatoridae == s; !ok {
		t.Fatalf("alligatoridae != %s", s)
	}
}
