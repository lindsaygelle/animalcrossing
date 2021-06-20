package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMRResettiAnimal(t *testing.T) {
	var s string = animals.Mole.Name()
	if ok := MRResetti.Animal() == s; !ok {
		t.Fatalf("%s != %s", MRResetti.Animal(), s)
	}
}

func TestMRResettiName(t *testing.T) {
	if ok := MRResetti.Name() == mrResetti; !ok {
		t.Fatalf("%s != %s", MRResetti.Name(), mrResetti)
	}
}
