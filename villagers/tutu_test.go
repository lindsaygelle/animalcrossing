package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestTutuName(t *testing.T) {
	if ok := Tutu.Name() == tutu; !ok {
		t.Fatalf("%s != %s", Tutu.Name(), tutu)
	}
}

func TestTutuSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Tutu.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tutu.Animal(), s)
	}
}
