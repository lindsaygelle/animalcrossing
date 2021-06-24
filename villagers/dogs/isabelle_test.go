package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestIsabelleAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Isabelle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Isabelle.Animal(), s)
	}
}

func TestIsabelleName(t *testing.T) {
	if ok := Isabelle.Name() == isabelle; !ok {
		t.Fatalf("%s != %s", Isabelle.Name(), isabelle)
	}
}
