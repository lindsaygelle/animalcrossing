package families

import "testing"

func TestFamilyAmbystomatidae(t *testing.T) {
	var s string = "Ambystomatidae"
	if ok := ambystomatidae == s; !ok {
		t.Fatalf("ambystomatidae != %s", s)
	}
}
