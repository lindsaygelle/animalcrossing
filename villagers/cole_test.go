package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestColeAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Cole.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cole.Animal(), s)
	}
}

func TestColeName(t *testing.T) {
	if ok := Cole.Name() == cole; !ok {
		t.Fatalf("%s != %s", Cole.Name(), cole)
	}
}
