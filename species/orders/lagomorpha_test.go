package orders

import "testing"

func TestLagomorpha(t *testing.T) {
	var s string = "Lagomorpha"
	if ok := lagomorpha == s; !ok {
		t.Fatalf("lagomorpha != %s", s)
	}
}
