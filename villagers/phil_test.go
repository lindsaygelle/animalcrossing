package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPhilAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Phil.Animal() == s; !ok {
		t.Fatalf("%s != %s", Phil.Animal(), s)
	}
}

func TestPhilName(t *testing.T) {
	if ok := Phil.Name() == phil; !ok {
		t.Fatalf("%s != %s", Phil.Name(), phil)
	}
}
