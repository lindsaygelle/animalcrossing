package families

import "testing"

func TestFamilyMuridae(t *testing.T) {
	var s string = "Muridae"
	if ok := muridae == s; !ok {
		t.Fatalf("muridae != %s", s)
	}
}
