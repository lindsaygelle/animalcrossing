package monkeys

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEliseAnimal(t *testing.T) {
	var s string = animals.Monkey.Name()
	if ok := Elise.Animal() == s; !ok {
		t.Fatalf("%s != %s", Elise.Animal(), s)
	}
}

func TestEliseName(t *testing.T) {
	if ok := Elise.Name() == elise; !ok {
		t.Fatalf("%s != %s", Elise.Name(), elise)
	}
}
