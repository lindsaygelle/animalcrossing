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
	var s string = species.Alligator.Name()
	if ok := Drago.Animal() == s; !ok {
		t.Fatalf("%s != %s", Drago.Animal(), s)
	}
}
