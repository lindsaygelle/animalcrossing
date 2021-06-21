package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFrancineAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Francine.Animal() == s; !ok {
		t.Fatalf("%s != %s", Francine.Animal(), s)
	}
}

func TestFrancineName(t *testing.T) {
	if ok := Francine.Name() == francine; !ok {
		t.Fatalf("%s != %s", Francine.Name(), francine)
	}
}
