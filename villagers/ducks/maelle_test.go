package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMaelleAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Maelle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Maelle.Animal(), s)
	}
}

func TestMaelleName(t *testing.T) {
	if ok := Maelle.Name() == maelle; !ok {
		t.Fatalf("%s != %s", Maelle.Name(), maelle)
	}
}
