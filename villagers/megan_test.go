package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMeganName(t *testing.T) {
	if ok := Megan.Name() == megan; !ok {
		t.Fatalf("%s != %s", Megan.Name(), megan)
	}
}

func TestMeganSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Megan.Animal() == s; !ok {
		t.Fatalf("%s != %s", Megan.Animal(), s)
	}
}
