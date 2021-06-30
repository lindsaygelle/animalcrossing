package orders

import "testing"

func TestGalliformes(t *testing.T) {
	var s string = "Galliformes"
	if ok := galliformes == s; !ok {
		t.Fatalf("galliformes != %s", s)
	}
}
