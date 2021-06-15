package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPekoeName(t *testing.T) {
	if ok := Pekoe.Name() == pekoe; !ok {
		t.Fatalf("%s != %s", Pekoe.Name(), pekoe)
	}
}

func TestPekoeSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Pekoe.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pekoe.Animal(), s)
	}
}
