package families

import "testing"

func TestFamilyHominidae(t *testing.T) {
	var s string = "Hominidae"
	if ok := hominidae == s; !ok {
		t.Fatalf("i != %s", s)
	}
}
