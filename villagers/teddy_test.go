package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestTeddyName(t *testing.T) {
	if ok := Teddy.Name() == teddy; !ok {
		t.Fatalf("%s != %s", Teddy.Name(), teddy)
	}
}

func TestTeddySpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Teddy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Teddy.Animal(), s)
	}
}
