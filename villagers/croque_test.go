package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCroqueAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Croque.Animal() == s; !ok {
		t.Fatalf("%s != %s", Croque.Animal(), s)
	}
}

func TestCroqueName(t *testing.T) {
	if ok := Croque.Name() == croque; !ok {
		t.Fatalf("%s != %s", Croque.Name(), croque)
	}
}
