package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestQueenieAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Queenie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Queenie.Animal(), s)
	}
}

func TestQueenieName(t *testing.T) {
	if ok := Queenie.Name() == queenie; !ok {
		t.Fatalf("%s != %s", Queenie.Name(), queenie)
	}
}
