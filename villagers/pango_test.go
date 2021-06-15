package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPangoName(t *testing.T) {
	if ok := Pango.Name() == pango; !ok {
		t.Fatalf("%s != %s", Pango.Name(), pango)
	}
}

func TestPangoSpecies(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Pango.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pango.Animal(), s)
	}
}
