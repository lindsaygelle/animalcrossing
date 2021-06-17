package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEloiseAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Eloise.Animal() == s; !ok {
		t.Fatalf("%s != %s", Eloise.Animal(), s)
	}
}

func TestEloiseName(t *testing.T) {
	if ok := Eloise.Name() == eloise; !ok {
		t.Fatalf("%s != %s", Eloise.Name(), eloise)
	}
}
