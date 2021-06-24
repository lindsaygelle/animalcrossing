package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMollyAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Molly.Animal() == s; !ok {
		t.Fatalf("%s != %s", Molly.Animal(), s)
	}
}

func TestMollyName(t *testing.T) {
	if ok := Molly.Name() == molly; !ok {
		t.Fatalf("%s != %s", Molly.Name(), molly)
	}
}
