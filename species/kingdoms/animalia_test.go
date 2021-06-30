package kingdoms

import "testing"

func TestKingdomAnimalia(t *testing.T) {
	var s string = "Animalia"
	if ok := animalia == s; !ok {
		t.Fatalf("animalia != %s", s)
	}
}
