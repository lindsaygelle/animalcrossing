package families

import "testing"

func TestFamilyAnatidae(t *testing.T) {
	var s string = "Anatidae"
	if ok := anatidae == s; !ok {
		t.Fatalf("anatidae != %s", s)
	}
}
