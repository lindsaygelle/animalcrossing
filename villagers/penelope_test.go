package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPenelopeAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Penelope.Animal() == s; !ok {
		t.Fatalf("%s != %s", Penelope.Animal(), s)
	}
}

func TestPenelopeName(t *testing.T) {
	if ok := Penelope.Name() == penelope; !ok {
		t.Fatalf("%s != %s", Penelope.Name(), penelope)
	}
}
