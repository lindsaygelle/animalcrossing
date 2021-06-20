package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGalaAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Gala.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gala.Animal(), s)
	}
}

func TestGalaName(t *testing.T) {
	if ok := Gala.Name() == gala; !ok {
		t.Fatalf("%s != %s", Gala.Name(), gala)
	}
}
