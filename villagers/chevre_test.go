package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChevreAnimal(t *testing.T) {
	var s string = animals.Goat.Name()
	if ok := Chevre.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chevre.Animal(), s)
	}
}

func TestChevreName(t *testing.T) {
	if ok := Chevre.Name() == chevre; !ok {
		t.Fatalf("%s != %s", Chevre.Name(), chevre)
	}
}
