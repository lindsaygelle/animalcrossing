package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestKodyName(t *testing.T) {
	if ok := Kody.Name() == kody; !ok {
		t.Fatalf("%s != %s", Kody.Name(), kody)
	}
}

func TestKodySpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Kody.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kody.Animal(), s)
	}
}
