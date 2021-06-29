package genuses

import "testing"

func TestGenusGallus(t *testing.T) {
	var s string = "Gallus"
	if ok := gallus == s; !ok {
		t.Fatalf("gallus != %s", s)
	}
}
