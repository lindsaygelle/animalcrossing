package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRickyAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Ricky.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ricky.Animal(), s)
	}
}

func TestRickyName(t *testing.T) {
	if ok := Ricky.Name() == ricky; !ok {
		t.Fatalf("%s != %s", Ricky.Name(), ricky)
	}
}
