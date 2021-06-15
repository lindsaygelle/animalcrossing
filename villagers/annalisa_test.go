package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAnnalisaName(t *testing.T) {
	if ok := Annalisa.Name() == annalisa; !ok {
		t.Fatalf("%s != %s", Annalisa.Name(), annalisa)
	}
}

func TestAnnalisaSpecies(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Annalisa.Animal() == s; !ok {
		t.Fatalf("%s != %s", Annalisa.Animal(), s)
	}
}
