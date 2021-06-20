package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPancettiAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Pancetti.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pancetti.Animal(), s)
	}
}

func TestPancettiName(t *testing.T) {
	if ok := Pancetti.Name() == pancetti; !ok {
		t.Fatalf("%s != %s", Pancetti.Name(), pancetti)
	}
}
