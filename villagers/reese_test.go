package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestReeseName(t *testing.T) {
	if ok := Reese.Name() == reese; !ok {
		t.Fatalf("%s != %s", Reese.Name(), reese)
	}
}

func TestReeseSpecies(t *testing.T) {
	var s string = species.Alpaca.Name()
	if ok := Reese.Animal() == s; !ok {
		t.Fatalf("%s != %s", Reese.Animal(), s)
	}
}
