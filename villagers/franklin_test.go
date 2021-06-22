package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFranklinAnimal(t *testing.T) {
	var s string = animals.Turkey.Name()
	if ok := Franklin.Animal() == s; !ok {
		t.Fatalf("%s != %s", Franklin.Animal(), s)
	}
}

func TestFranklinName(t *testing.T) {
	if ok := Franklin.Name() == franklin; !ok {
		t.Fatalf("%s != %s", Franklin.Name(), franklin)
	}
}
