package octopuses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestInkwellAnimal(t *testing.T) {
	var s string = animals.Octopus.Name()
	if ok := Inkwell.Animal() == s; !ok {
		t.Fatalf("%s != %s", Inkwell.Animal(), s)
	}
}

func TestInkwellName(t *testing.T) {
	if ok := Inkwell.Name() == inkwell; !ok {
		t.Fatalf("%s != %s", Inkwell.Name(), inkwell)
	}
}
