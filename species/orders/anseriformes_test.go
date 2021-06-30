package orders

import "testing"

func TestAnseriformes(t *testing.T) {
	var s string = "Anseriformes"
	if ok := anseriformes == s; !ok {
		t.Fatalf("anseriformes != %s", s)
	}
}
