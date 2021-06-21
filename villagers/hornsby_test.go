package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHornsbyAnimal(t *testing.T) {
	var s string = animals.Rhinoceros.Name()
	if ok := Hornsby.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hornsby.Animal(), s)
	}
}

func TestHornsbyName(t *testing.T) {
	if ok := Hornsby.Name() == hornsby; !ok {
		t.Fatalf("%s != %s", Hornsby.Name(), hornsby)
	}
}
