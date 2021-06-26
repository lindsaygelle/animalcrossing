package rabbits

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPippyAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Pippy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pippy.Animal(), s)
	}
}

func TestPippyName(t *testing.T) {
	if ok := Pippy.Name() == pippy; !ok {
		t.Fatalf("%s != %s", Pippy.Name(), pippy)
	}
}
