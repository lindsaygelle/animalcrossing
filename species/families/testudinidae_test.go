package families

import "testing"

func TestTestudinidae(t *testing.T) {
	var s string = "Testudinidae"
	if ok := testudinidae == s; !ok {
		t.Fatalf("testudinidae != %s", s)
	}
}
