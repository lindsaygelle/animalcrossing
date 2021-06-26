package gorillas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBooneAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Boone.Animal() == s; !ok {
		t.Fatalf("%s != %s", Boone.Animal(), s)
	}
}

func TestBooneName(t *testing.T) {
	if ok := Boone.Name() == boone; !ok {
		t.Fatalf("%s != %s", Boone.Name(), boone)
	}
}
