package families

import "testing"

func TestAnatidae(t *testing.T) {
	var s string = "Anatidae"
	if ok := anatidae == s; !ok {
		t.Fatalf("anatidae != %s", s)
	}
}
