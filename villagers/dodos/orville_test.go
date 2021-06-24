package dodos

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOrvilleAnimal(t *testing.T) {
	var s string = animals.Dodo.Name()
	if ok := Orville.Animal() == s; !ok {
		t.Fatalf("%s != %s", Orville.Animal(), s)
	}
}

func TestOrvilleName(t *testing.T) {
	if ok := Orville.Name() == orville; !ok {
		t.Fatalf("%s != %s", Orville.Name(), orville)
	}
}
