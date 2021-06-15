package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAlliName(t *testing.T) {
	if ok := Alli.Name() == alli; !ok {
		t.Fatalf("%s != %s", Alli.Name(), alli)
	}
}

func TestAlliSpecies(t *testing.T) {
	var s string = animals.Alligator.Name()
	if ok := Alli.Animal() == s; !ok {
		t.Fatalf("%s != %s", Alli.Animal(), s)
	}
}
