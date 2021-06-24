package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRoverAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Rover.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rover.Animal(), s)
	}
}

func TestRoverName(t *testing.T) {
	if ok := Rover.Name() == rover; !ok {
		t.Fatalf("%s != %s", Rover.Name(), rover)
	}
}
