package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestGrizzlyName(t *testing.T) {
	if ok := Grizzly.Name() == grizzly; !ok {
		t.Fatalf("%s != %s", Grizzly.Name(), grizzly)
	}
}

func TestGrizzlySpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Grizzly.Animal() == s; !ok {
		t.Fatalf("%s != %s", Grizzly.Animal(), s)
	}
}
