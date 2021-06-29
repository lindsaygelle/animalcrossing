package families

import "testing"

func TestMuridae(t *testing.T) {
	var s string = "Muridae"
	if ok := muridae == s; !ok {
		t.Fatalf("muridae != %s", s)
	}
}
