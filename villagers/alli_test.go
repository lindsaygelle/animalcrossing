package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestAlliName(t *testing.T) {
	if ok := Alli.Name() == alli; !ok {
		t.Fatalf("%s != %s", Alli.Name(), alli)
	}
}

func TestAlliSpecies(t *testing.T) {
	if ok := Alli.Animal() == species.Alligator.Name(); !ok {
		t.Fatalf("%s != %s", Alli.Animal(), species.Alligator.Name())
	}
}
