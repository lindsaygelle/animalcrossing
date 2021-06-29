package families

import "testing"

func TestPelecanidae(t *testing.T) {
	var s string = "Pelecanidae"
	if ok := pelecanidae == s; !ok {
		t.Fatalf("pelecanidae != %s", s)
	}
}
