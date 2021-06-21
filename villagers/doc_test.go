package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDocAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Doc.Animal() == s; !ok {
		t.Fatalf("%s != %s", Doc.Animal(), s)
	}
}

func TestDocName(t *testing.T) {
	if ok := Doc.Name() == doc; !ok {
		t.Fatalf("%s != %s", Doc.Name(), doc)
	}
}
