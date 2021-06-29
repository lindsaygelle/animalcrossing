package families

import "testing"

func TestFamilyTalpidae(t *testing.T) {
	var s string = "Talpidae"
	if ok := talpidae == s; !ok {
		t.Fatalf("talpidae != %s", talpidae)
	}
}
