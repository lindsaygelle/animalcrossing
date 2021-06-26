package mice

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChadderAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Chadder.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chadder.Animal(), s)
	}
}

func TestChadderName(t *testing.T) {
	if ok := Chadder.Name() == chadder; !ok {
		t.Fatalf("%s != %s", Chadder.Name(), chadder)
	}
}
