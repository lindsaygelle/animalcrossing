package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestDozerName(t *testing.T) {
	if ok := Dozer.Name() == dozer; !ok {
		t.Fatalf("%s != %s", Dozer.Name(), dozer)
	}
}

func TestDozerSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Dozer.Animal() == s; !ok {
		t.Fatalf("%s != %s", Dozer.Animal(), s)
	}
}
