package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPeteAnimal(t *testing.T) {
	var s string = animals.Pelican.Name()
	if ok := Pete.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pete.Animal(), s)
	}
}

func TestPeteName(t *testing.T) {
	if ok := Pete.Name() == pete; !ok {
		t.Fatalf("%s != %s", Pete.Name(), pete)
	}
}
