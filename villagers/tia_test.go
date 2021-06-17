package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTiaAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Tia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tia.Animal(), s)
	}
}

func TestTiaName(t *testing.T) {
	if ok := Tia.Name() == tia; !ok {
		t.Fatalf("%s != %s", Tia.Name(), tia)
	}
}
