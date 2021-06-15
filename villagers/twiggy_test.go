package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTwiggyName(t *testing.T) {
	if ok := Twiggy.Name() == twiggy; !ok {
		t.Fatalf("%s != %s", Twiggy.Name(), twiggy)
	}
}

func TestTwiggySpecies(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Twiggy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Twiggy.Animal(), s)
	}
}
