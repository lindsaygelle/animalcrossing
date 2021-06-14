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
	if ok := Gayle.Animal() == species.Alligator.Name(); !ok {
		t.Fatalf("%s != %s", Gayle.Animal(), species.Alligator.Name())
	}
}
