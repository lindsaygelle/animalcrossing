package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAisleName(t *testing.T) {
	if ok := Aisle.Name() == aisle; !ok {
		t.Fatalf("%s != %s", Aisle.Name(), aisle)
	}
}

func TestAisleSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Aisle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Aisle.Animal(), s)
	}
}
