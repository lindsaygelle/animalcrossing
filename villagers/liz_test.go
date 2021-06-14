package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestLizName(t *testing.T) {
	if ok := Liz.Name() == liz; !ok {
		t.Fatalf("%s != %s", Liz.Name(), liz)
	}
}

func TestLizSpecies(t *testing.T) {
	if ok := Liz.Animal() == species.Alligator.Name(); !ok {
		t.Fatalf("%s != %s", Liz.Animal(), species.Alligator.Name())
	}
}
