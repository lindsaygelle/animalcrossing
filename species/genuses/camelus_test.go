package genuses

import "testing"

func TestCamelus(t *testing.T) {
	var s string = "Camelus"
	if ok := camelus == s; !ok {
		t.Fatalf("camelus != %s", s)
	}
}
