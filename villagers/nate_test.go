package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestNateName(t *testing.T) {
	if ok := Nate.Name() == nate; !ok {
		t.Fatalf("%s != %s", Nate.Name(), nate)
	}
}

func TestNateSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Nate.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nate.Animal(), s)
	}
}
