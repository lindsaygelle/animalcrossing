package hamsters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHamphreyAnimal(t *testing.T) {
	var s string = animals.Hamster.Name()
	if ok := Hamphrey.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hamphrey.Animal(), s)
	}
}

func TestHamphreyName(t *testing.T) {
	if ok := Hamphrey.Name() == hamphrey; !ok {
		t.Fatalf("%s != %s", Hamphrey.Name(), hamphrey)
	}
}
