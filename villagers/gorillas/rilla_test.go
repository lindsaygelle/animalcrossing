package gorillas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRillaAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Rilla.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rilla.Animal(), s)
	}
}

func TestRillaName(t *testing.T) {
	if ok := Rilla.Name() == rilla; !ok {
		t.Fatalf("%s != %s", Rilla.Name(), rilla)
	}
}
