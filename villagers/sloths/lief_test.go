package sloths

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLiefAnimal(t *testing.T) {
	var s string = animals.Sloth.Name()
	if ok := Lief.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lief.Animal(), s)
	}
}

func TestLiefName(t *testing.T) {
	if ok := Lief.Name() == lief; !ok {
		t.Fatalf("%s != %s", Lief.Name(), lief)
	}
}
