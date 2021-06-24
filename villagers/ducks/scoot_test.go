package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestScootAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Scoot.Animal() == s; !ok {
		t.Fatalf("%s != %s", Scoot.Animal(), s)
	}
}

func TestScootName(t *testing.T) {
	if ok := Scoot.Name() == scoot; !ok {
		t.Fatalf("%s != %s", Scoot.Name(), scoot)
	}
}
