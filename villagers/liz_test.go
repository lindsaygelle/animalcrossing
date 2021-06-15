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
	var s string = species.Alligator.Name()
	if ok := Liz.Animal() == s; !ok {
		t.Fatalf("%s != %s", Liz.Animal(), s)
	}
}
