package pigs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGanonAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Ganon.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ganon.Animal(), s)
	}
}

func TestGanonName(t *testing.T) {
	if ok := Ganon.Name() == ganon; !ok {
		t.Fatalf("%s != %s", Ganon.Name(), ganon)
	}
}
