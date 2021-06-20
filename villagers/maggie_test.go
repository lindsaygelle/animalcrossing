package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMaggieAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Maggie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Maggie.Animal(), s)
	}
}

func TestMaggieName(t *testing.T) {
	if ok := Maggie.Name() == maggie; !ok {
		t.Fatalf("%s != %s", Maggie.Name(), maggie)
	}
}
