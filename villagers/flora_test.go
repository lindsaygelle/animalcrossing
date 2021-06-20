package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFloraAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Flora.Animal() == s; !ok {
		t.Fatalf("%s != %s", Flora.Animal(), s)
	}
}

func TestFloraName(t *testing.T) {
	if ok := Flora.Name() == flora; !ok {
		t.Fatalf("%s != %s", Flora.Name(), flora)
	}
}
