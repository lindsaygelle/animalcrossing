package families

import "testing"

func TestFamilyCricetidae(t *testing.T) {
	var s string = "Cricetidae"
	if ok := cricetidae == s; !ok {
		t.Fatalf("cricetidae != %s", s)
	}
}
