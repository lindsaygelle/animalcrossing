package species

import "testing"

func TestFCatus(t *testing.T) {
	var s string = "F. catus"
	if ok := fCatus == s; !ok {
		t.Fatalf("fCatus != %s", s)
	}
}
