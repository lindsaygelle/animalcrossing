package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFilbertAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Filbert.Animal() == s; !ok {
		t.Fatalf("%s != %s", Filbert.Animal(), s)
	}
}

func TestFilbertName(t *testing.T) {
	if ok := Filbert.Name() == filbert; !ok {
		t.Fatalf("%s != %s", Filbert.Name(), filbert)
	}
}
