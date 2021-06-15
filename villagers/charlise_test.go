package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCharliseAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Charlise.Animal() == s; !ok {
		t.Fatalf("%s != %s", Charlise.Animal(), s)
	}
}

func TestCharliseName(t *testing.T) {
	if ok := Charlise.Name() == charlise; !ok {
		t.Fatalf("%s != %s", Charlise.Name(), charlise)
	}
}
