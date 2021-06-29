package genuses

import "testing"

func TestGenusVicugna(t *testing.T) {
	var s string = "Vicugna"
	if ok := vicugna == s; !ok {
		t.Fatalf("vicugna != %s", s)
	}
}
