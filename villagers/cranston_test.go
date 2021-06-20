package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCranstonAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Cranston.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cranston.Animal(), s)
	}
}

func TestCranstonName(t *testing.T) {
	if ok := Cranston.Name() == cranston; !ok {
		t.Fatalf("%s != %s", Cranston.Name(), cranston)
	}
}
