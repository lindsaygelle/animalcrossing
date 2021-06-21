package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTommyAnimal(t *testing.T) {
	var s string = animals.Raccoon.Name()
	if ok := Tommy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tommy.Animal(), s)
	}
}

func TestTommyName(t *testing.T) {
	if ok := Tommy.Name() == tommy; !ok {
		t.Fatalf("%s != %s", Tommy.Name(), tommy)
	}
}
