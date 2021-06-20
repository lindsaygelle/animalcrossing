package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHughAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Hugh.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hugh.Animal(), s)
	}
}

func TestHughName(t *testing.T) {
	if ok := Hugh.Name() == hugh; !ok {
		t.Fatalf("%s != %s", Hugh.Name(), hugh)
	}
}
