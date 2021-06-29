package families

import "testing"

func TestFamilyEquidae(t *testing.T) {
	var s string = "Equidae"
	if ok := equidae == s; !ok {
		t.Fatalf("equidae != %s", s)
	}
}
