package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSydneyAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Sydney.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sydney.Animal(), s)
	}
}

func TestSydneyName(t *testing.T) {
	if ok := Sydney.Name() == sydney; !ok {
		t.Fatalf("%s != %s", Sydney.Name(), sydney)
	}
}
