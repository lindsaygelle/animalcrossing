package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOliveName(t *testing.T) {
	if ok := Olive.Name() == olive; !ok {
		t.Fatalf("%s != %s", Olive.Name(), olive)
	}
}

func TestOliveSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Olive.Animal() == s; !ok {
		t.Fatalf("%s != %s", Olive.Animal(), s)
	}
}
