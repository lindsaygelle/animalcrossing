package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestStinkyAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Stinky.Animal() == s; !ok {
		t.Fatalf("%s != %s", Stinky.Animal(), s)
	}
}

func TestStinkyName(t *testing.T) {
	if ok := Stinky.Name() == stinky; !ok {
		t.Fatalf("%s != %s", Stinky.Name(), stinky)
	}
}
