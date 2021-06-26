package ostriches

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBlancheAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Blanche.Animal() == s; !ok {
		t.Fatalf("%s != %s", Blanche.Animal(), s)
	}
}

func TestBlancheName(t *testing.T) {
	if ok := Blanche.Name() == blanche; !ok {
		t.Fatalf("%s != %s", Blanche.Name(), blanche)
	}
}
