package genuses

import "testing"

func TestOvis(t *testing.T) {
	var s string = "Ovis"
	if ok := ovis == s; !ok {
		t.Fatalf("ovis != %s", s)
	}
}
