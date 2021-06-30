package orders

import "testing"

func TestCaudata(t *testing.T) {
	var s string = "Caudata"
	if ok := caudata == s; !ok {
		t.Fatalf("caudata != %s", s)
	}
}
