package orders

import "testing"

func TestSphenisciformes(t *testing.T) {
	var s string = "Sphenisciformes"
	if ok := sphenisciformes == s; !ok {
		t.Fatalf("sphenisciformes != %s", s)
	}
}
