package orders

import "testing"

func TestAccipitriformes(t *testing.T) {
	var s string = "Accipitriformes"
	if ok := accipitriformes == s; !ok {
		t.Fatalf("accipitriformes != %s", s)
	}
}
