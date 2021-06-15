package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGayleName(t *testing.T) {
	if ok := Gayle.Name() == gayle; !ok {
		t.Fatalf("%s != %s", Gayle.Name(), gayle)
	}
}

func TestGayleSpecies(t *testing.T) {
	var s string = animals.Alligator.Name()
	if ok := Gayle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gayle.Animal(), s)
	}
}
