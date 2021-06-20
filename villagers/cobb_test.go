package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCobbAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Cobb.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cobb.Animal(), s)
	}
}

func TestCobbName(t *testing.T) {
	if ok := Cobb.Name() == cobb; !ok {
		t.Fatalf("%s != %s", Cobb.Name(), cobb)
	}
}
