package families

import "testing"

func TestMustelidae(t *testing.T) {
	var s string = "Mustelidae"
	if ok := mustelidae == s; !ok {
		t.Fatalf("mustelidae != %s", s)
	}
}
