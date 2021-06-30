package species

import "testing"

func TestOAries(t *testing.T) {
	var s string = "O. aries"
	if ok := oAries == s; !ok {
		t.Fatalf("oAries != %s", s)
	}
}
