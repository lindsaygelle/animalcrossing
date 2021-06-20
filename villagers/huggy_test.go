package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHuggyAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Huggy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Huggy.Animal(), s)
	}
}

func TestHuggyName(t *testing.T) {
	if ok := Huggy.Name() == huggy; !ok {
		t.Fatalf("%s != %s", Huggy.Name(), huggy)
	}
}
