package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBowAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Bow.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bow.Animal(), s)
	}
}

func TestBowName(t *testing.T) {
	if ok := Bow.Name() == bow; !ok {
		t.Fatalf("%s != %s", Bow.Name(), bow)
	}
}
