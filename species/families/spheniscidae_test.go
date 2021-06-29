package families

import "testing"

func TestFamilySpheniscidae(t *testing.T) {
	var s string = "Spheniscidae"
	if ok := spheniscidae == s; !ok {
		t.Fatalf("spheniscidae != %s", s)
	}
}
