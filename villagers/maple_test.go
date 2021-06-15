package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestMapleName(t *testing.T) {
	if ok := Maple.Name() == maple; !ok {
		t.Fatalf("%s != %s", Maple.Name(), maple)
	}
}

func TestMapleSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Maple.Animal() == s; !ok {
		t.Fatalf("%s != %s", Maple.Animal(), s)
	}
}
