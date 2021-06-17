package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMaddieAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Maddie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Maddie.Animal(), s)
	}
}

func TestMaddieName(t *testing.T) {
	if ok := Maddie.Name() == maddie; !ok {
		t.Fatalf("%s != %s", Maddie.Name(), maddie)
	}
}
