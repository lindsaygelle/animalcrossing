package monkeys

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFlipAnimal(t *testing.T) {
	var s string = animals.Monkey.Name()
	if ok := Flip.Animal() == s; !ok {
		t.Fatalf("%s != %s", Flip.Animal(), s)
	}
}

func TestFlipName(t *testing.T) {
	if ok := Flip.Name() == flip; !ok {
		t.Fatalf("%s != %s", Flip.Name(), flip)
	}
}
