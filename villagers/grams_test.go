package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGramsAnimal(t *testing.T) {
	var s string = animals.Kappa.Name()
	if ok := Grams.Animal() == s; !ok {
		t.Fatalf("%s != %s", Grams.Animal(), s)
	}
}

func TestGramsName(t *testing.T) {
	if ok := Grams.Name() == grams; !ok {
		t.Fatalf("%s != %s", Grams.Name(), grams)
	}
}
