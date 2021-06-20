package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMooseAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Moose.Animal() == s; !ok {
		t.Fatalf("%s != %s", Moose.Animal(), s)
	}
}

func TestMooseName(t *testing.T) {
	if ok := Moose.Name() == moose; !ok {
		t.Fatalf("%s != %s", Moose.Name(), moose)
	}
}
