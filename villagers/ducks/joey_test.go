package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJoeyAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Joey.Animal() == s; !ok {
		t.Fatalf("%s != %s", Joey.Animal(), s)
	}
}

func TestJoeyName(t *testing.T) {
	if ok := Joey.Name() == joey; !ok {
		t.Fatalf("%s != %s", Joey.Name(), joey)
	}
}
