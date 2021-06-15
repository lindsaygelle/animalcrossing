package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestCharliseName(t *testing.T) {
	if ok := Charlise.Name() == charlise; !ok {
		t.Fatalf("%s != %s", Charlise.Name(), charlise)
	}
}

func TestCharliseSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Charlise.Animal() == s; !ok {
		t.Fatalf("%s != %s", Charlise.Animal(), s)
	}
}
