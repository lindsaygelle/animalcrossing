package genuses

import "testing"

func TestGenusAmbystoma(t *testing.T) {
	var s string = "Ambystoma"
	if ok := ambystoma == s; !ok {
		t.Fatalf("ambystoma != %s", s)
	}
}
