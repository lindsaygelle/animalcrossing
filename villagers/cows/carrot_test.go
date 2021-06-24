package cows

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCarrotAnimal(t *testing.T) {
	var s string = animals.Cow.Name()
	if ok := Carrot.Animal() == s; !ok {
		t.Fatalf("%s != %s", Carrot.Animal(), s)
	}
}

func TestCarrotName(t *testing.T) {
	if ok := Carrot.Name() == carrot; !ok {
		t.Fatalf("%s != %s", Carrot.Name(), carrot)
	}
}
