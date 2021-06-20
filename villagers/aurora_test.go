package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAuroraAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Aurora.Animal() == s; !ok {
		t.Fatalf("%s != %s", Aurora.Animal(), s)
	}
}

func TestAuroraName(t *testing.T) {
	if ok := Aurora.Name() == aurora; !ok {
		t.Fatalf("%s != %s", Aurora.Name(), aurora)
	}
}
