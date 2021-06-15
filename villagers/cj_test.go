package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestCJName(t *testing.T) {
	if ok := CJ.Name() == cj; !ok {
		t.Fatalf("%s != %s", CJ.Name(), cj)
	}
}

func TestCJSpecies(t *testing.T) {
	var s string = species.Beaver.Name()
	if ok := CJ.Animal() == s; !ok {
		t.Fatalf("%s != %s", CJ.Animal(), s)
	}
}
