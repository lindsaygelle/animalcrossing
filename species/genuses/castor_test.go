package genuses

import "testing"

func TestGenusCastor(t *testing.T) {
	var s string = "Castor"
	if ok := castor == s; !ok {
		t.Fatalf("castor != %s", s)
	}
}
