package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestDelName(t *testing.T) {
	if ok := Del.Name() == del; !ok {
		t.Fatalf("%s != %s", Del.Name(), del)
	}
}

func TestDelSpecies(t *testing.T) {
	if ok := Del.Animal() == species.Alligator.Name(); !ok {
		t.Fatalf("%s != %s", Del.Animal(), species.Alligator.Name())
	}
}
