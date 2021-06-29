package genuses

import "testing"

func TestGenusOdobenus(t *testing.T) {
	var s string = "Odobenus"
	if ok := odobenus == s; !ok {
		t.Fatalf("odobenus != %s", s)
	}
}
