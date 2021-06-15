package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAnchovyName(t *testing.T) {
	if ok := Anchovy.Name() == anchovy; !ok {
		t.Fatalf("%s != %s", Anchovy.Name(), anchovy)
	}
}

func TestAnchovySpecies(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Anchovy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Anchovy.Animal(), s)
	}
}
