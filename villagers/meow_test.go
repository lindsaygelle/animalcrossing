package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMeowAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Meow.Animal() == s; !ok {
		t.Fatalf("%s != %s", Meow.Animal(), s)
	}
}

func TestMeowName(t *testing.T) {
	if ok := Meow.Name() == meow; !ok {
		t.Fatalf("%s != %s", Meow.Name(), meow)
	}
}
