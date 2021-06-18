package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestVelmaAnimal(t *testing.T) {
	var s string = animals.Goat.Name()
	if ok := Velma.Animal() == s; !ok {
		t.Fatalf("%s != %s", Velma.Animal(), s)
	}
}

func TestVelmaName(t *testing.T) {
	if ok := Velma.Name() == velma; !ok {
		t.Fatalf("%s != %s", Velma.Name(), velma)
	}
}
