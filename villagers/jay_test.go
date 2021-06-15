package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJayName(t *testing.T) {
	if ok := Jay.Name() == jay; !ok {
		t.Fatalf("%s != %s", Jay.Name(), jay)
	}
}

func TestJaySpecies(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Jay.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jay.Animal(), s)
	}
}
