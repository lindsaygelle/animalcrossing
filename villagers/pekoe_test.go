package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestPekoeName(t *testing.T) {
	if ok := Pekoe.Name() == pekoe; !ok {
		t.Fatalf("%s != %s", Pekoe.Name(), pekoe)
	}
}

func TestPekoeSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Pekoe.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pekoe.Animal(), s)
	}
}
