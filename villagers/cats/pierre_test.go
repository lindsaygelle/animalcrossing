package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPierreAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Pierre.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pierre.Animal(), s)
	}
}

func TestPierreName(t *testing.T) {
	if ok := Pierre.Name() == pierre; !ok {
		t.Fatalf("%s != %s", Pierre.Name(), pierre)
	}
}
