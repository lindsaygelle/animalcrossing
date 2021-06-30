package species

import "testing"

func TestGCamelopardalis(t *testing.T) {
	var s string = "G. camelopardalis"
	if ok := gCamelopardalis == s; !ok {
		t.Fatalf("gCamelopardalis != %s", s)
	}
}
