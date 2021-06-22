package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCeceAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Cece.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cece.Animal(), s)
	}
}

func TestCeceName(t *testing.T) {
	if ok := Cece.Name() == cece; !ok {
		t.Fatalf("%s != %s", Cece.Name(), cece)
	}
}
