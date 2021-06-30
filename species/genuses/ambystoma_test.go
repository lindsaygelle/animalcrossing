package genuses

import "testing"

func TestAmbystoma(t *testing.T) {
	var s string = "Ambystoma"
	if ok := ambystoma == s; !ok {
		t.Fatalf("ambystoma != %s", s)
	}
}
