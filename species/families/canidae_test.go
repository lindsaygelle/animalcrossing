package families

import "testing"

func TestFamilyCanidae(t *testing.T) {
	var s string = "Canidae"
	if ok := canidae == s; !ok {
		t.Fatalf("canidae != %s", s)
	}
}
