package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRoaldAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Roald.Animal() == s; !ok {
		t.Fatalf("%s != %s", Roald.Animal(), s)
	}
}

func TestRoaldName(t *testing.T) {
	if ok := Roald.Name() == roald; !ok {
		t.Fatalf("%s != %s", Roald.Name(), roald)
	}
}
