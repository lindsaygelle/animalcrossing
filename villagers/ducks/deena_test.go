package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDeenaAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Deena.Animal() == s; !ok {
		t.Fatalf("%s != %s", Deena.Animal(), s)
	}
}

func TestDeenaName(t *testing.T) {
	if ok := Deena.Name() == deena; !ok {
		t.Fatalf("%s != %s", Deena.Name(), deena)
	}
}
