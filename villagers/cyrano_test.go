package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestCyranoName(t *testing.T) {
	if ok := Cyrano.Name() == cyrano; !ok {
		t.Fatalf("%s != %s", Cyrano.Name(), cyrano)
	}
}

func TestCyranoSpecies(t *testing.T) {
	var s string = species.Anteater.Name()
	if ok := Cyrano.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cyrano.Animal(), s)
	}
}
