package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChipAnimal(t *testing.T) {
	var s string = animals.Beaver.Name()
	if ok := Chip.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chip.Animal(), s)
	}
}

func TestChipName(t *testing.T) {
	if ok := Chip.Name() == chip; !ok {
		t.Fatalf("%s != %s", Chip.Name(), chip)
	}
}
