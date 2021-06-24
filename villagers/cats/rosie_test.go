package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRosieAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Rosie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rosie.Animal(), s)
	}
}

func TestRosieName(t *testing.T) {
	if ok := Rosie.Name() == rosie; !ok {
		t.Fatalf("%s != %s", Rosie.Name(), rosie)
	}
}
