package orders

import "testing"

func TestColumbiformes(t *testing.T) {
	var s string = "Columbiformes"
	if ok := columbiformes == s; !ok {
		t.Fatalf("columbiformes != %s", s)
	}
}
