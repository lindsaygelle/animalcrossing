package genuses

import "testing"

func TestStruthio(t *testing.T) {
	var s string = "Struthio"
	if ok := struthio == s; !ok {
		t.Fatalf("struthio != %s", s)
	}
}
