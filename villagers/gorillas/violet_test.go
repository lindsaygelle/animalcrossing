package gorillas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestVioletAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Violet.Animal() == s; !ok {
		t.Fatalf("%s != %s", Violet.Animal(), s)
	}
}

func TestVioletName(t *testing.T) {
	if ok := Violet.Name() == violet; !ok {
		t.Fatalf("%s != %s", Violet.Name(), violet)
	}
}
