package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDragoAnimal(t *testing.T) {
	var s string = animals.Alligator.Name()
	if ok := Drago.Animal() == s; !ok {
		t.Fatalf("%s != %s", Drago.Animal(), s)
	}
}

func TestDragoName(t *testing.T) {
	if ok := Drago.Name() == drago; !ok {
		t.Fatalf("%s != %s", Drago.Name(), drago)
	}
}
