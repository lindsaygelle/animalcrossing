package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFrobertAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Frobert.Animal() == s; !ok {
		t.Fatalf("%s != %s", Frobert.Animal(), s)
	}
}

func TestFrobertName(t *testing.T) {
	if ok := Frobert.Name() == frobert; !ok {
		t.Fatalf("%s != %s", Frobert.Name(), frobert)
	}
}
