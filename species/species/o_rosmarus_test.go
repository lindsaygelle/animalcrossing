package species

import "testing"

func TestORosmarus(t *testing.T) {
	var s string = "O. rosmarus"
	if ok := oRosmarus == s; !ok {
		t.Fatalf("oRosmarus != %s", s)
	}
}
