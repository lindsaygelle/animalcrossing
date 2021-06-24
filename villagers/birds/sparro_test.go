package birds

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSparroAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Sparro.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sparro.Animal(), s)
	}
}

func TestSparroName(t *testing.T) {
	if ok := Sparro.Name() == sparro; !ok {
		t.Fatalf("%s != %s", Sparro.Name(), sparro)
	}
}
