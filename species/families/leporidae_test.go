package families

import "testing"

func TestFamilyLeporidae(t *testing.T) {
	var s string = "Leporidae"
	if ok := leporidae == s; !ok {
		t.Fatalf("leporidae != %s", s)
	}
}
