package orders

import "testing"

func TestCalumbiformes(t *testing.T) {
	var s string = "Calumbiformes"
	if ok := calumbiformes == s; !ok {
		t.Fatalf("calumbiformes != %s", s)
	}
}
