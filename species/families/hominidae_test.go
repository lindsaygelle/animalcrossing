package families

import "testing"

func TestHominidae(t *testing.T) {
	var s string = "Hominidae"
	if ok := hominidae == s; !ok {
		t.Fatalf("i != %s", s)
	}
}
