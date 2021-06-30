package orders

import "testing"

func TestPelecaniformes(t *testing.T) {
	var s string = "Pelecaniformes"
	if ok := pelecaniformes == s; !ok {
		t.Fatalf("pelecaniformes != %s", s)
	}
}
