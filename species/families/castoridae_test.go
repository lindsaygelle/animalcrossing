package families

import "testing"

func TestFamilyCastoridae(t *testing.T) {
	var s string = "Castoridae"
	if ok := castoridae == s; !ok {
		t.Fatalf("castoridae != %s", s)
	}
}
