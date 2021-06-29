package families

import "testing"

func TestFamilyCervidae(t *testing.T) {
	var s string = "Cervidae"
	if ok := cervidae == s; !ok {
		t.Fatalf("cervidae != %s", s)
	}
}
