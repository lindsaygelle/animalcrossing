package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestIkeName(t *testing.T) {
	if ok := Ike.Name() == ike; !ok {
		t.Fatalf("%s != %s", Ike.Name(), ike)
	}
}

func TestIkeSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Ike.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ike.Animal(), s)
	}
}
