package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestGrouchoName(t *testing.T) {
	if ok := Groucho.Name() == groucho; !ok {
		t.Fatalf("%s != %s", Groucho.Name(), groucho)
	}
}

func TestGrouchoSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Groucho.Animal() == s; !ok {
		t.Fatalf("%s != %s", Groucho.Animal(), s)
	}
}
