package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMargieAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Margie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Margie.Animal(), s)
	}
}

func TestMargieName(t *testing.T) {
	if ok := Margie.Name() == margie; !ok {
		t.Fatalf("%s != %s", Margie.Name(), margie)
	}
}
