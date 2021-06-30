package orders

import "testing"

func TestCarnivora(t *testing.T) {
	var s string = "Carnivora"
	if ok := carnivora == s; !ok {
		t.Fatalf("carnivora != %s", s)
	}
}
