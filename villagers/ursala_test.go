package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestUrsalaName(t *testing.T) {
	if ok := Ursala.Name() == ursala; !ok {
		t.Fatalf("%s != %s", Ursala.Name(), ursala)
	}
}

func TestUrsalaSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Ursala.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ursala.Animal(), s)
	}
}
