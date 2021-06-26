package hamsters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFlurryAnimal(t *testing.T) {
	var s string = animals.Hamster.Name()
	if ok := Flurry.Animal() == s; !ok {
		t.Fatalf("%s != %s", Flurry.Animal(), s)
	}
}

func TestFlurryName(t *testing.T) {
	if ok := Flurry.Name() == flurry; !ok {
		t.Fatalf("%s != %s", Flurry.Name(), flurry)
	}
}
