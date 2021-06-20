package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBrewsterAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Brewster.Animal() == s; !ok {
		t.Fatalf("%s != %s", Brewster.Animal(), s)
	}
}

func TestBrewsterName(t *testing.T) {
	if ok := Brewster.Name() == brewster; !ok {
		t.Fatalf("%s != %s", Brewster.Name(), brewster)
	}
}
