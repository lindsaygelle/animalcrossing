package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGooseAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Goose.Animal() == s; !ok {
		t.Fatalf("%s != %s", Goose.Animal(), s)
	}
}

func TestGooseName(t *testing.T) {
	if ok := Goose.Name() == goose; !ok {
		t.Fatalf("%s != %s", Goose.Name(), goose)
	}
}
