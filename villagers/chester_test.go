package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestChesterName(t *testing.T) {
	if ok := Chester.Name() == chester; !ok {
		t.Fatalf("%s != %s", Chester.Name(), chester)
	}
}

func TestChesterSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Chester.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chester.Animal(), s)
	}
}
