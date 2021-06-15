package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestChipName(t *testing.T) {
	if ok := Chip.Name() == chip; !ok {
		t.Fatalf("%s != %s", Chip.Name(), chip)
	}
}

func TestChipSpecies(t *testing.T) {
	var s string = species.Beaver.Name()
	if ok := Chip.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chip.Animal(), s)
	}
}
