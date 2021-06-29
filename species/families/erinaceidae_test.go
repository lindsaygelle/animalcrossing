package families

import "testing"

func TestFamilyErinaceidae(t *testing.T) {
	var s string = "Erinaceidae"
	if ok := erinaceidae == s; !ok {
		t.Fatalf("erinaceidae != %s", s)
	}
}
