package owls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCelesteAnimal(t *testing.T) {
	var s string = animals.Owl.Name()
	if ok := Celeste.Animal() == s; !ok {
		t.Fatalf("%s != %s", Celeste.Animal(), s)
	}
}

func TestCelesteName(t *testing.T) {
	if ok := Celeste.Name() == celeste; !ok {
		t.Fatalf("%s != %s", Celeste.Name(), celeste)
	}
}
