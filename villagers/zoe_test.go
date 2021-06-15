package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestZoeName(t *testing.T) {
	if ok := Zoe.Name() == zoe; !ok {
		t.Fatalf("%s != %s", Zoe.Name(), zoe)
	}
}

func TestZoeSpecies(t *testing.T) {
	var s string = species.Anteater.Name()
	if ok := Zoe.Animal() == s; !ok {
		t.Fatalf("%s != %s", Zoe.Animal(), s)
	}
}
