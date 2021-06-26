package hamsters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGrahamAnimal(t *testing.T) {
	var s string = animals.Hamster.Name()
	if ok := Graham.Animal() == s; !ok {
		t.Fatalf("%s != %s", Graham.Animal(), s)
	}
}

func TestGrahamName(t *testing.T) {
	if ok := Graham.Name() == graham; !ok {
		t.Fatalf("%s != %s", Graham.Name(), graham)
	}
}
