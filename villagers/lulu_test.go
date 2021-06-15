package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestLuluName(t *testing.T) {
	if ok := Lulu.Name() == lulu; !ok {
		t.Fatalf("%s != %s", Lulu.Name(), lulu)
	}
}

func TestLuluSpecies(t *testing.T) {
	var s string = species.Anteater.Name()
	if ok := Lulu.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lulu.Animal(), s)
	}
}
