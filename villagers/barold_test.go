package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBaroldName(t *testing.T) {
	if ok := Barold.Name() == barold; !ok {
		t.Fatalf("%s != %s", Barold.Name(), barold)
	}
}

func TestBaroldSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Barold.Animal() == s; !ok {
		t.Fatalf("%s != %s", Barold.Animal(), s)
	}
}
