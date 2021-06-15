package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestPinkyName(t *testing.T) {
	if ok := Pinky.Name() == pinky; !ok {
		t.Fatalf("%s != %s", Pinky.Name(), pinky)
	}
}

func TestPinkySpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Pinky.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pinky.Animal(), s)
	}
}
