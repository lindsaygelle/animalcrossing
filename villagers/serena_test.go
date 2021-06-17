package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSerenaAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Serena.Animal() == s; !ok {
		t.Fatalf("%s != %s", Serena.Animal(), s)
	}
}

func TestSerenaName(t *testing.T) {
	if ok := Serena.Name() == serena; !ok {
		t.Fatalf("%s != %s", Serena.Name(), serena)
	}
}
