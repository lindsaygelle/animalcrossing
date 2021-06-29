package families

import "testing"

func TestFamilyAgamidae(t *testing.T) {
	var s string = "Agamidae"
	if ok := agamidae == s; !ok {
		t.Fatalf("agamidae != %s", s)
	}
}
