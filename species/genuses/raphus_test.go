package genuses

import "testing"

func TestGenusRaphus(t *testing.T) {
	var s string = "Raphus"
	if ok := raphus == s; !ok {
		t.Fatalf("raphus != %s", s)
	}
}
