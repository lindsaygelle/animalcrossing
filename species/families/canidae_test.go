package families

import "testing"

func TestCanidae(t *testing.T) {
	var s string = "Canidae"
	if ok := canidae == s; !ok {
		t.Fatalf("canidae != %s", s)
	}
}
