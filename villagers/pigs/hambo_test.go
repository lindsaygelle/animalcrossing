package pigs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHamboAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Hambo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hambo.Animal(), s)
	}
}

func TestHamboName(t *testing.T) {
	if ok := Hambo.Name() == hambo; !ok {
		t.Fatalf("%s != %s", Hambo.Name(), hambo)
	}
}
