package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPecanAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Pecan.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pecan.Animal(), s)
	}
}

func TestPecanName(t *testing.T) {
	if ok := Pecan.Name() == pecan; !ok {
		t.Fatalf("%s != %s", Pecan.Name(), pecan)
	}
}
