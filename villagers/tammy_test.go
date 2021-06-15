package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/species"
)

func TestTammyName(t *testing.T) {
	if ok := Tammy.Name() == tammy; !ok {
		t.Fatalf("%s != %s", Tammy.Name(), tammy)
	}
}

func TestTammySpecies(t *testing.T) {
	var s string = species.Bear.Name()
	if ok := Tammy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tammy.Animal(), s)
	}
}
