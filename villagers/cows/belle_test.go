package cows

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBelleAnimal(t *testing.T) {
	var s string = animals.Cow.Name()
	if ok := Belle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Belle.Animal(), s)
	}
}

func TestBelleName(t *testing.T) {
	if ok := Belle.Name() == belle; !ok {
		t.Fatalf("%s != %s", Belle.Name(), belle)
	}
}
