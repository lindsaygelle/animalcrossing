package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestJuneName(t *testing.T) {
	if ok := June.Name() == june; !ok {
		t.Fatalf("%s != %s", June.Name(), june)
	}
}

func TestJuneSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := June.Animal() == s; !ok {
		t.Fatalf("%s != %s", June.Animal(), s)
	}
}
