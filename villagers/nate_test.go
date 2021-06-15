package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNateName(t *testing.T) {
	if ok := Nate.Name() == nate; !ok {
		t.Fatalf("%s != %s", Nate.Name(), nate)
	}
}

func TestNateSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Nate.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nate.Animal(), s)
	}
}
