package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHopperAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Hopper.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hopper.Animal(), s)
	}
}

func TestHopperName(t *testing.T) {
	if ok := Hopper.Name() == hopper; !ok {
		t.Fatalf("%s != %s", Hopper.Name(), hopper)
	}
}
