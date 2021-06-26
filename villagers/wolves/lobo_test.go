package wolves

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLoboAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Lobo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lobo.Animal(), s)
	}
}

func TestLoboName(t *testing.T) {
	if ok := Lobo.Name() == lobo; !ok {
		t.Fatalf("%s != %s", Lobo.Name(), lobo)
	}
}
