package rabbits

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBunnieAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Bunnie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bunnie.Animal(), s)
	}
}

func TestBunnieName(t *testing.T) {
	if ok := Bunnie.Name() == bunnie; !ok {
		t.Fatalf("%s != %s", Bunnie.Name(), bunnie)
	}
}
