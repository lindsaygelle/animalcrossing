package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestDragoName(t *testing.T) {
	if ok := Drago.Name() == drago; !ok {
		t.Fatalf("%s != %s", Drago.Name(), drago)
	}
}

func TestDragoSpecies(t *testing.T) {
	if ok := Drago.Animal() == species.Alligator.Name(); !ok {
		t.Fatalf("%s != %s", Drago.Animal(), species.Alligator.Name())
	}
}
