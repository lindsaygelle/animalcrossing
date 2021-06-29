package families

import "testing"

func TestErinaceidae(t *testing.T) {
	var s string = "Erinaceidae"
	if ok := erinaceidae == s; !ok {
		t.Fatalf("erinaceidae != %s", s)
	}
}
