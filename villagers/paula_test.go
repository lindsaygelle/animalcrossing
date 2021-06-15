package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPaulaName(t *testing.T) {
	if ok := Paula.Name() == paula; !ok {
		t.Fatalf("%s != %s", Paula.Name(), paula)
	}
}

func TestPaulaSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Paula.Animal() == s; !ok {
		t.Fatalf("%s != %s", Paula.Animal(), s)
	}
}
