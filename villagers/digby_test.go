package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDigbyAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Digby.Animal() == s; !ok {
		t.Fatalf("%s != %s", Digby.Animal(), s)
	}
}

func TestDigbyName(t *testing.T) {
	if ok := Digby.Name() == digby; !ok {
		t.Fatalf("%s != %s", Digby.Name(), digby)
	}
}
