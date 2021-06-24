package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKetchupAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Ketchup.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ketchup.Animal(), s)
	}
}

func TestKetchupName(t *testing.T) {
	if ok := Ketchup.Name() == ketchup; !ok {
		t.Fatalf("%s != %s", Ketchup.Name(), ketchup)
	}
}
