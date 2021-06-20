package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWadeAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Wade.Animal() == s; !ok {
		t.Fatalf("%s != %s", Wade.Animal(), s)
	}
}

func TestWadeName(t *testing.T) {
	if ok := Wade.Name() == wade; !ok {
		t.Fatalf("%s != %s", Wade.Name(), wade)
	}
}
