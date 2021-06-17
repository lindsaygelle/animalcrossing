package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBiskitAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Biskit.Animal() == s; !ok {
		t.Fatalf("%s != %s", Biskit.Animal(), s)
	}
}

func TestBiskitName(t *testing.T) {
	if ok := Biskit.Name() == biskit; !ok {
		t.Fatalf("%s != %s", Biskit.Name(), biskit)
	}
}
