package families

import "testing"

func TestAlligatoridae(t *testing.T) {
	var s string = "Alligatoridae"
	if ok := alligatoridae == s; !ok {
		t.Fatalf("alligatoridae != %s", s)
	}
}
