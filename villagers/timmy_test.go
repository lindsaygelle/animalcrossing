package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTimmyAnimal(t *testing.T) {
	var s string = animals.Raccoon.Name()
	if ok := Timmy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Timmy.Animal(), s)
	}
}

func TestTimmyName(t *testing.T) {
	if ok := Timmy.Name() == timmy; !ok {
		t.Fatalf("%s != %s", Timmy.Name(), timmy)
	}
}
