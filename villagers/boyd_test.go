package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBoydAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Boyd.Animal() == s; !ok {
		t.Fatalf("%s != %s", Boyd.Animal(), s)
	}
}

func TestBoydName(t *testing.T) {
	if ok := Boyd.Name() == boyd; !ok {
		t.Fatalf("%s != %s", Boyd.Name(), boyd)
	}
}
