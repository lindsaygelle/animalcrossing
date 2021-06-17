package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMarcelAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Marcel.Animal() == s; !ok {
		t.Fatalf("%s != %s", Marcel.Animal(), s)
	}
}

func TestMarcelName(t *testing.T) {
	if ok := Marcel.Name() == marcel; !ok {
		t.Fatalf("%s != %s", Marcel.Name(), marcel)
	}
}
