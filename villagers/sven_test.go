package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSvenAnimal(t *testing.T) {
	var s string = animals.Goat.Name()
	if ok := Sven.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sven.Animal(), s)
	}
}

func TestSvenName(t *testing.T) {
	if ok := Sven.Name() == sven; !ok {
		t.Fatalf("%s != %s", Sven.Name(), sven)
	}
}
