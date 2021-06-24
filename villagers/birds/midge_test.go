package birds

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMidgeAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Midge.Animal() == s; !ok {
		t.Fatalf("%s != %s", Midge.Animal(), s)
	}
}

func TestMidgeName(t *testing.T) {
	if ok := Midge.Name() == midge; !ok {
		t.Fatalf("%s != %s", Midge.Name(), midge)
	}
}
