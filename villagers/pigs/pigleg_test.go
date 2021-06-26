package pigs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPiglegAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Pigleg.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pigleg.Animal(), s)
	}
}

func TestPiglegName(t *testing.T) {
	if ok := Pigleg.Name() == pigleg; !ok {
		t.Fatalf("%s != %s", Pigleg.Name(), pigleg)
	}
}
