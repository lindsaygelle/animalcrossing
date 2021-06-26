package pigs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCurlyAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Curly.Animal() == s; !ok {
		t.Fatalf("%s != %s", Curly.Animal(), s)
	}
}

func TestCurlyName(t *testing.T) {
	if ok := Curly.Name() == curly; !ok {
		t.Fatalf("%s != %s", Curly.Name(), curly)
	}
}
