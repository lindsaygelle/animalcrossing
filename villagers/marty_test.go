package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestMartyName(t *testing.T) {
	if ok := Marty.Name() == marty; !ok {
		t.Fatalf("%s != %s", Marty.Name(), marty)
	}
}

func TestMartySpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Marty.Animal() == s; !ok {
		t.Fatalf("%s != %s", Marty.Animal(), s)
	}
}
