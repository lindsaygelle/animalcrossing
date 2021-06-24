package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHenryAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Henry.Animal() == s; !ok {
		t.Fatalf("%s != %s", Henry.Animal(), s)
	}
}

func TestHenryName(t *testing.T) {
	if ok := Henry.Name() == henry; !ok {
		t.Fatalf("%s != %s", Henry.Name(), henry)
	}
}
