package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMallaryAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Mallary.Animal() == s; !ok {
		t.Fatalf("%s != %s", Mallary.Animal(), s)
	}
}

func TestMallaryName(t *testing.T) {
	if ok := Mallary.Name() == mallary; !ok {
		t.Fatalf("%s != %s", Mallary.Name(), mallary)
	}
}
