package families

import "testing"

func TestSpheniscidae(t *testing.T) {
	var s string = "Spheniscidae"
	if ok := spheniscidae == s; !ok {
		t.Fatalf("spheniscidae != %s", s)
	}
}
