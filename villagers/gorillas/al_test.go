package gorillas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAlAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Al.Animal() == s; !ok {
		t.Fatalf("%s != %s", Al.Animal(), s)
	}
}

func TestAlName(t *testing.T) {
	if ok := Al.Name() == al; !ok {
		t.Fatalf("%s != %s", Al.Name(), al)
	}
}
