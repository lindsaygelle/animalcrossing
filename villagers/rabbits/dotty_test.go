package rabbits

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDottyAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Dotty.Animal() == s; !ok {
		t.Fatalf("%s != %s", Dotty.Animal(), s)
	}
}

func TestDottyName(t *testing.T) {
	if ok := Dotty.Name() == dotty; !ok {
		t.Fatalf("%s != %s", Dotty.Name(), dotty)
	}
}
