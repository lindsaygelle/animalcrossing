package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestVivianAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Vivian.Animal() == s; !ok {
		t.Fatalf("%s != %s", Vivian.Animal(), s)
	}
}

func TestVivianName(t *testing.T) {
	if ok := Vivian.Name() == vivian; !ok {
		t.Fatalf("%s != %s", Vivian.Name(), vivian)
	}
}
