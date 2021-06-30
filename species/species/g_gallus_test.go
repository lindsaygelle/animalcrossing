package species

import "testing"

func TestGGallus(t *testing.T) {
	var s string = "G. gallus"
	if ok := gGallus == s; !ok {
		t.Fatalf("gGallus != %s", s)
	}
}
