package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCJName(t *testing.T) {
	if ok := CJ.Name() == cj; !ok {
		t.Fatalf("%s != %s", CJ.Name(), cj)
	}
}

func TestCJSpecies(t *testing.T) {
	var s string = animals.Beaver.Name()
	if ok := CJ.Animal() == s; !ok {
		t.Fatalf("%s != %s", CJ.Animal(), s)
	}
}
