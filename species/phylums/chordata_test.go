package phylums

import "testing"

func TestPhylumChordata(t *testing.T) {
	var s string = "Chordata"
	if ok := chordata == s; !ok {
		t.Fatalf("chordata != %s", s)
	}
}
