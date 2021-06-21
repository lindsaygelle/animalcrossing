package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTobyAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Toby.Animal() == s; !ok {
		t.Fatalf("%s != %s", Toby.Animal(), s)
	}
}

func TestTobyName(t *testing.T) {
	if ok := Toby.Name() == toby; !ok {
		t.Fatalf("%s != %s", Toby.Name(), toby)
	}
}
