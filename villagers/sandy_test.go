package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSandyAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Sandy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sandy.Animal(), s)
	}
}

func TestSandyName(t *testing.T) {
	if ok := Sandy.Name() == sandy; !ok {
		t.Fatalf("%s != %s", Sandy.Name(), sandy)
	}
}
