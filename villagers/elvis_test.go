package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestElvisAnimal(t *testing.T) {
	var s string = animals.Lion.Name()
	if ok := Elvis.Animal() == s; !ok {
		t.Fatalf("%s != %s", Elvis.Animal(), s)
	}
}

func TestElvisName(t *testing.T) {
	if ok := Elvis.Name() == elvis; !ok {
		t.Fatalf("%s != %s", Elvis.Name(), elvis)
	}
}
