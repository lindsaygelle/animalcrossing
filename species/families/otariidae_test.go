package families

import "testing"

func TestFamilyOtariidae(t *testing.T) {
	var s string = "Otariidae"
	if ok := otariidae == s; !ok {
		t.Fatalf("otariidae != %s", s)
	}
}
