package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestVicheAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Viche.Animal() == s; !ok {
		t.Fatalf("%s != %s", Viche.Animal(), s)
	}
}

func TestVicheName(t *testing.T) {
	if ok := Viche.Name() == viche; !ok {
		t.Fatalf("%s != %s", Viche.Name(), viche)
	}
}
