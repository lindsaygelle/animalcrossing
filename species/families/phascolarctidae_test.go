package families

import "testing"

func TestPhascolarctidae(t *testing.T) {
	var s string = "Phascolarctidae"
	if ok := phascolarctidae == s; !ok {
		t.Fatalf("phascolarctidae != %s", s)
	}
}
