package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCupcakeAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Cupcake.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cupcake.Animal(), s)
	}
}

func TestCupcakeName(t *testing.T) {
	if ok := Cupcake.Name() == cupcake; !ok {
		t.Fatalf("%s != %s", Cupcake.Name(), cupcake)
	}
}
