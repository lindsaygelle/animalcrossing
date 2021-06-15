package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestOlafName(t *testing.T) {
	if ok := Olaf.Name() == olaf; !ok {
		t.Fatalf("%s != %s", Olaf.Name(), olaf)
	}
}

func TestOlafSpecies(t *testing.T) {
	var s string = species.Anteater.Name()
	if ok := Olaf.Animal() == s; !ok {
		t.Fatalf("%s != %s", Olaf.Animal(), s)
	}
}
