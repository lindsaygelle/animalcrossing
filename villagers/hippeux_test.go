package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHippeuxAnimal(t *testing.T) {
	var s string = animals.Hippo.Name()
	if ok := Hippeux.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hippeux.Animal(), s)
	}
}

func TestHippeuxName(t *testing.T) {
	if ok := Hippeux.Name() == hippeux; !ok {
		t.Fatalf("%s != %s", Hippeux.Name(), hippeux)
	}
}
