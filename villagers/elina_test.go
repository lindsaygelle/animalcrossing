package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestElinaAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Elina.Animal() == s; !ok {
		t.Fatalf("%s != %s", Elina.Animal(), s)
	}
}

func TestElinaName(t *testing.T) {
	if ok := Elina.Name() == elina; !ok {
		t.Fatalf("%s != %s", Elina.Name(), elina)
	}
}
