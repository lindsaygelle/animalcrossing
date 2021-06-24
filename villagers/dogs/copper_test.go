package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCopperAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Copper.Animal() == s; !ok {
		t.Fatalf("%s != %s", Copper.Animal(), s)
	}
}

func TestCopperName(t *testing.T) {
	if ok := Copper.Name() == copper; !ok {
		t.Fatalf("%s != %s", Copper.Name(), copper)
	}
}
