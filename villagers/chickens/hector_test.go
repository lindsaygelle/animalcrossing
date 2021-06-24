package chickens

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHectorAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Hector.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hector.Animal(), s)
	}
}

func TestHectorName(t *testing.T) {
	if ok := Hector.Name() == hector; !ok {
		t.Fatalf("%s != %s", Hector.Name(), hector)
	}
}
