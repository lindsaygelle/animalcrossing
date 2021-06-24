package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGloriaAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Gloria.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gloria.Animal(), s)
	}
}

func TestGloriaName(t *testing.T) {
	if ok := Gloria.Name() == gloria; !ok {
		t.Fatalf("%s != %s", Gloria.Name(), gloria)
	}
}
