package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestReneeAnimal(t *testing.T) {
	var s string = animals.Rhinoceros.Name()
	if ok := Renee.Animal() == s; !ok {
		t.Fatalf("%s != %s", Renee.Animal(), s)
	}
}

func TestReneeName(t *testing.T) {
	if ok := Renee.Name() == renee; !ok {
		t.Fatalf("%s != %s", Renee.Name(), renee)
	}
}
