package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestNosegayName(t *testing.T) {
	if ok := Nosegay.Name() == nosegay; !ok {
		t.Fatalf("%s != %s", Nosegay.Name(), nosegay)
	}
}

func TestNosegaySpecies(t *testing.T) {
	var s string = species.Anteater.Name()
	if ok := Nosegay.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nosegay.Animal(), s)
	}
}
