package alligators

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSlyAnimal(t *testing.T) {
	var s string = animals.Alligator.Name()
	if ok := Sly.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sly.Animal(), s)
	}
}

func TestSlyName(t *testing.T) {
	if ok := Sly.Name() == sly; !ok {
		t.Fatalf("%s != %s", Sly.Name(), sly)
	}
}
