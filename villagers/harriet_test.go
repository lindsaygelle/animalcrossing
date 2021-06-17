package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHarrietAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Harriet.Animal() == s; !ok {
		t.Fatalf("%s != %s", Harriet.Animal(), s)
	}
}

func TestHarrietName(t *testing.T) {
	if ok := Harriet.Name() == harriet; !ok {
		t.Fatalf("%s != %s", Harriet.Name(), harriet)
	}
}
