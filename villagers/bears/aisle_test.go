package bears

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAisleAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Aisle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Aisle.Animal(), s)
	}
}

func TestAisleName(t *testing.T) {
	if ok := Aisle.Name() == aisle; !ok {
		t.Fatalf("%s != %s", Aisle.Name(), aisle)
	}
}
