package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCurtName(t *testing.T) {
	if ok := Curt.Name() == curt; !ok {
		t.Fatalf("%s != %s", Curt.Name(), curt)
	}
}

func TestCurtSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Curt.Animal() == s; !ok {
		t.Fatalf("%s != %s", Curt.Animal(), s)
	}
}
