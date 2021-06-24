package chickens

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPluckyAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Plucky.Animal() == s; !ok {
		t.Fatalf("%s != %s", Plucky.Animal(), s)
	}
}

func TestPluckyName(t *testing.T) {
	if ok := Plucky.Name() == plucky; !ok {
		t.Fatalf("%s != %s", Plucky.Name(), plucky)
	}
}
