package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMoniqueAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Monique.Animal() == s; !ok {
		t.Fatalf("%s != %s", Monique.Animal(), s)
	}
}

func TestMoniqueName(t *testing.T) {
	if ok := Monique.Name() == monique; !ok {
		t.Fatalf("%s != %s", Monique.Name(), monique)
	}
}
