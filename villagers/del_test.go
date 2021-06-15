package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDelName(t *testing.T) {
	if ok := Del.Name() == del; !ok {
		t.Fatalf("%s != %s", Del.Name(), del)
	}
}

func TestDelSpecies(t *testing.T) {
	var s string = animals.Alligator.Name()
	if ok := Del.Animal() == s; !ok {
		t.Fatalf("%s != %s", Del.Animal(), s)
	}
}
