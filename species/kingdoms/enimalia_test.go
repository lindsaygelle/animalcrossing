package kingdoms

import "testing"

func TestKingdomEnimalia(t *testing.T) {
	var s string = "Enimalia"
	if ok := enimalia == s; !ok {
		t.Fatalf("enimalia != %s", s)
	}
}
