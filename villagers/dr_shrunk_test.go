package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDrShrunkName(t *testing.T) {
	if ok := DrShrunk.Name() == drShrunk; !ok {
		t.Fatalf("%s != %s", DrShrunk.Name(), drShrunk)
	}
}

func TestDrShrunkSpecies(t *testing.T) {
	var s string = animals.Axolotl.Name()
	if ok := DrShrunk.Animal() == s; !ok {
		t.Fatalf("%s != %s", DrShrunk.Animal(), s)
	}
}
