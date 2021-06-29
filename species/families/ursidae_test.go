package families

import "testing"

func TestFamilyUrsidae(t *testing.T) {
	var s string = "Ursidae"
	if ok := ursidae == s; !ok {
		t.Fatalf("ursidae != %s", s)
	}
}
