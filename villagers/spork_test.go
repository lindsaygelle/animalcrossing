package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSporkAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Spork.Animal() == s; !ok {
		t.Fatalf("%s != %s", Spork.Animal(), s)
	}
}

func TestSporkName(t *testing.T) {
	if ok := Spork.Name() == spork; !ok {
		t.Fatalf("%s != %s", Spork.Name(), spork)
	}
}
