package species

import "testing"

func TestCKingii(t *testing.T) {
	var s string = "C. kingii"
	if ok := cKingii == s; !ok {
		t.Fatalf("cKingii != %s", s)
	}
}
