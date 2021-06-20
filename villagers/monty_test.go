package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMontyAnimal(t *testing.T) {
	var s string = animals.Monkey.Name()
	if ok := Monty.Animal() == s; !ok {
		t.Fatalf("%s != %s", Monty.Animal(), s)
	}
}

func TestMontyName(t *testing.T) {
	if ok := Monty.Name() == monty; !ok {
		t.Fatalf("%s != %s", Monty.Name(), monty)
	}
}
