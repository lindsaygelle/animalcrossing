package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPateAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Pate.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pate.Animal(), s)
	}
}

func TestPateName(t *testing.T) {
	if ok := Pate.Name() == pate; !ok {
		t.Fatalf("%s != %s", Pate.Name(), pate)
	}
}
