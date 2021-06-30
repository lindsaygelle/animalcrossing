package genuses

import "testing"

func TestOdobenus(t *testing.T) {
	var s string = "Odobenus"
	if ok := odobenus == s; !ok {
		t.Fatalf("odobenus != %s", s)
	}
}
