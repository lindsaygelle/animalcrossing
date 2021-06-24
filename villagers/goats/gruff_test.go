package goats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGruffAnimal(t *testing.T) {
	var s string = animals.Goat.Name()
	if ok := Gruff.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gruff.Animal(), s)
	}
}

func TestGruffName(t *testing.T) {
	if ok := Gruff.Name() == gruff; !ok {
		t.Fatalf("%s != %s", Gruff.Name(), gruff)
	}
}
