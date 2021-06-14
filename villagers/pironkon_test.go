package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestPironkonName(t *testing.T) {
	if ok := Pironkon.Name() == pironkon; !ok {
		t.Fatalf("%s != %s", Pironkon.Name(), pironkon)
	}
}

func TestPironkonSpecies(t *testing.T) {
	if ok := Pironkon.Animal() == species.Alligator.Name(); !ok {
		t.Fatalf("%s != %s", Pironkon.Animal(), species.Alligator.Name())
	}
}
