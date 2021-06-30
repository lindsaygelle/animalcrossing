package genuses

import "testing"

func TestVicugna(t *testing.T) {
	var s string = "Vicugna"
	if ok := vicugna == s; !ok {
		t.Fatalf("vicugna != %s", s)
	}
}
