package animalcrossing

import "testing"

func TestPantherName(t *testing.T) {
	if ok := Panther.Name() == panther; !ok {
		t.Fatal("Panther.Name() != panther")
	}
}
