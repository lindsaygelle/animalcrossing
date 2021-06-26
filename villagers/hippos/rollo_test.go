package hippopotamuses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRolloAnimal(t *testing.T) {
	var s string = animals.Hippopotamus.Name()
	if ok := Rollo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rollo.Animal(), s)
	}
}

func TestRolloName(t *testing.T) {
	if ok := Rollo.Name() == rollo; !ok {
		t.Fatalf("%s != %s", Rollo.Name(), rollo)
	}
}
