package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestUrsalaName(t *testing.T) {
	if ok := Ursala.Name() == ursala; !ok {
		t.Fatalf("%s != %s", Ursala.Name(), ursala)
	}
}

func TestUrsalaSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Ursala.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ursala.Animal(), s)
	}
}
