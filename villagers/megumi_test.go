package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMegumiAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Megumi.Animal() == s; !ok {
		t.Fatalf("%s != %s", Megumi.Animal(), s)
	}
}

func TestMegumiName(t *testing.T) {
	if ok := Megumi.Name() == megumi; !ok {
		t.Fatalf("%s != %s", Megumi.Name(), megumi)
	}
}
