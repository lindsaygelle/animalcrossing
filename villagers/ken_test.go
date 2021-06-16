package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKenAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Ken.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ken.Animal(), s)
	}
}

func TestKenName(t *testing.T) {
	if ok := Ken.Name() == ken; !ok {
		t.Fatalf("%s != %s", Ken.Name(), ken)
	}
}
