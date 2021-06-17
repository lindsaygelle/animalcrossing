package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChampagneAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Champagne.Animal() == s; !ok {
		t.Fatalf("%s != %s", Champagne.Animal(), s)
	}
}

func TestChampagneName(t *testing.T) {
	if ok := Champagne.Name() == champagne; !ok {
		t.Fatalf("%s != %s", Champagne.Name(), champagne)
	}
}
