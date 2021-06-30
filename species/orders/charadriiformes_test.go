package orders

import "testing"

func TestCharadriiformes(t *testing.T) {
	var s string = "Charadriiformes"
	if ok := charadriiformes == s; !ok {
		t.Fatalf("charadriiformes != %s", s)
	}
}
