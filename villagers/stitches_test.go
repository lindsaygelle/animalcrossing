package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestStitchesName(t *testing.T) {
	if ok := Stitches.Name() == stitches; !ok {
		t.Fatalf("%s != %s", Stitches.Name(), stitches)
	}
}

func TestStitchesSpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Stitches.Animal() == s; !ok {
		t.Fatalf("%s != %s", Stitches.Animal(), s)
	}
}
