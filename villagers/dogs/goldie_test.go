package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGoldieAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Goldie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Goldie.Animal(), s)
	}
}

func TestGoldieName(t *testing.T) {
	if ok := Goldie.Name() == goldie; !ok {
		t.Fatalf("%s != %s", Goldie.Name(), goldie)
	}
}
