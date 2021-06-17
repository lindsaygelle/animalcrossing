package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTadAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Tad.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tad.Animal(), s)
	}
}

func TestTadName(t *testing.T) {
	if ok := Tad.Name() == tad; !ok {
		t.Fatalf("%s != %s", Tad.Name(), tad)
	}
}
