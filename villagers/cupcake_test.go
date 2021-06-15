package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestCupcakeName(t *testing.T) {
	if ok := Cupcake.Name() == cupcake; !ok {
		t.Fatalf("%s != %s", Cupcake.Name(), cupcake)
	}
}

func TestCupcakeSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Cupcake.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cupcake.Animal(), s)
	}
}
