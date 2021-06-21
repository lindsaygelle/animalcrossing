package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTankAnimal(t *testing.T) {
	var s string = animals.Rhinoceros.Name()
	if ok := Tank.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tank.Animal(), s)
	}
}

func TestTankName(t *testing.T) {
	if ok := Tank.Name() == tank; !ok {
		t.Fatalf("%s != %s", Tank.Name(), tank)
	}
}
