package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestGayleName(t *testing.T) {
	if ok := Gayle.Name() == gayle; !ok {
		t.Fatalf("%s != %s", Gayle.Name(), gayle)
	}
}

func TestGayleSpecies(t *testing.T) {
	var s string = species.Alligator.Name()
	if ok := Gayle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gayle.Animal(), s)
	}
}
