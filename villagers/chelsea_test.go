package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChelseaAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Chelsea.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chelsea.Animal(), s)
	}
}

func TestChelseaName(t *testing.T) {
	if ok := Chelsea.Name() == chelsea; !ok {
		t.Fatalf("%s != %s", Chelsea.Name(), chelsea)
	}
}
