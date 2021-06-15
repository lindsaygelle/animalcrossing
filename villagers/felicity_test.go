package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFelicityAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Felicity.Animal() == s; !ok {
		t.Fatalf("%s != %s", Felicity.Animal(), s)
	}
}

func TestFelicityName(t *testing.T) {
	if ok := Felicity.Name() == felicity; !ok {
		t.Fatalf("%s != %s", Felicity.Name(), felicity)
	}
}
