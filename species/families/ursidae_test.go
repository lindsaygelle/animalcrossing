package families

import "testing"

func TestUrsidae(t *testing.T) {
	var s string = "Ursidae"
	if ok := ursidae == s; !ok {
		t.Fatalf("ursidae != %s", s)
	}
}
