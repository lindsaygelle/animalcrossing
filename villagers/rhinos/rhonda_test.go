package rhinoceroses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRhondaAnimal(t *testing.T) {
	var s string = animals.Rhinoceros.Name()
	if ok := Rhonda.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rhonda.Animal(), s)
	}
}

func TestRhondaName(t *testing.T) {
	if ok := Rhonda.Name() == rhonda; !ok {
		t.Fatalf("%s != %s", Rhonda.Name(), rhonda)
	}
}
