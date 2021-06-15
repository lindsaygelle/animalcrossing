package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKlausName(t *testing.T) {
	if ok := Klaus.Name() == klaus; !ok {
		t.Fatalf("%s != %s", Klaus.Name(), klaus)
	}
}

func TestKlausSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Klaus.Animal() == s; !ok {
		t.Fatalf("%s != %s", Klaus.Animal(), s)
	}
}
