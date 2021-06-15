package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGiovanniName(t *testing.T) {
	if ok := Giovanni.Name() == giovanni; !ok {
		t.Fatalf("%s != %s", Giovanni.Name(), giovanni)
	}
}

func TestGiovanniSpecies(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Giovanni.Animal() == s; !ok {
		t.Fatalf("%s != %s", Giovanni.Animal(), s)
	}
}
