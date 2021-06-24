package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCousteauAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Cousteau.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cousteau.Animal(), s)
	}
}

func TestCousteauName(t *testing.T) {
	if ok := Cousteau.Name() == cousteau; !ok {
		t.Fatalf("%s != %s", Cousteau.Name(), cousteau)
	}
}
