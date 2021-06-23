package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWolfgangAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Wolfgang.Animal() == s; !ok {
		t.Fatalf("%s != %s", Wolfgang.Animal(), s)
	}
}

func TestWolfgangName(t *testing.T) {
	if ok := Wolfgang.Name() == wolfgang; !ok {
		t.Fatalf("%s != %s", Wolfgang.Name(), wolfgang)
	}
}
