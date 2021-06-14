package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestSlyName(t *testing.T) {
	if ok := Sly.Name() == sly; !ok {
		t.Fatalf("%s != %s", Sly.Name(), sly)
	}
}

func TestSlySpecies(t *testing.T) {
	if ok := Sly.Animal() == species.Alligator.Name(); !ok {
		t.Fatalf("%s != %s", Sly.Animal(), species.Alligator.Name())
	}
}
