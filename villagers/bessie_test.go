package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBessieAnimal(t *testing.T) {
	var s string = animals.Cow.Name()
	if ok := Bessie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bessie.Animal(), s)
	}
}

func TestBessieName(t *testing.T) {
	if ok := Bessie.Name() == bessie; !ok {
		t.Fatalf("%s != %s", Bessie.Name(), bessie)
	}
}
