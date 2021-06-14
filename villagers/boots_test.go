package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestBootsName(t *testing.T) {
	if ok := Boots.Name() == boots; !ok {
		t.Fatalf("%s != %s", Boots.Name(), boots)
	}
}

func TestBootsSpecies(t *testing.T) {
	if ok := Boots.Animal() == species.Alligator.Name(); !ok {
		t.Fatalf("%s != %s", Boots.Animal(), species.Alligator.Name())
	}
}
