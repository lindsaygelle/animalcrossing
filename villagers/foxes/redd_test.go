package foxes

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestReddAnimal(t *testing.T) {
	var s string = animals.Fox.Name()
	if ok := Redd.Animal() == s; !ok {
		t.Fatalf("%s != %s", Redd.Animal(), s)
	}
}

func TestReddName(t *testing.T) {
	if ok := Redd.Name() == redd; !ok {
		t.Fatalf("%s != %s", Redd.Name(), redd)
	}
}
