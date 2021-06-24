package goats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestIggyAnimal(t *testing.T) {
	var s string = animals.Goat.Name()
	if ok := Iggy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Iggy.Animal(), s)
	}
}

func TestIggyName(t *testing.T) {
	if ok := Iggy.Name() == iggy; !ok {
		t.Fatalf("%s != %s", Iggy.Name(), iggy)
	}
}
