package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestMurphyName(t *testing.T) {
	if ok := Murphy.Name() == murphy; !ok {
		t.Fatalf("%s != %s", Murphy.Name(), murphy)
	}
}

func TestMurphySpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Murphy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Murphy.Animal(), s)
	}
}
