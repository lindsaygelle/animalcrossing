package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestCyrusName(t *testing.T) {
	if ok := Cyrus.Name() == cyrus; !ok {
		t.Fatalf("%s != %s", Cyrus.Name(), cyrus)
	}
}

func TestCyrusSpecies(t *testing.T) {
	var s string = species.Alpaca.Name()
	if ok := Cyrus.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cyrus.Animal(), s)
	}
}
