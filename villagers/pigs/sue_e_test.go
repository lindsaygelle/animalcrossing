package pigs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSueEAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := SueE.Animal() == s; !ok {
		t.Fatalf("%s != %s", SueE.Animal(), s)
	}
}

func TestSueEName(t *testing.T) {
	if ok := SueE.Name() == sueE; !ok {
		t.Fatalf("%s != %s", SueE.Name(), sueE)
	}
}
