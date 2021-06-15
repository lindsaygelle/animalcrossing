package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestVladimirName(t *testing.T) {
	if ok := Vladimir.Name() == vladimir; !ok {
		t.Fatalf("%s != %s", Vladimir.Name(), vladimir)
	}
}

func TestVladimirSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Vladimir.Animal() == s; !ok {
		t.Fatalf("%s != %s", Vladimir.Animal(), s)
	}
}
