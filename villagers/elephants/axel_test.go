package elephants

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAxelAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Axel.Animal() == s; !ok {
		t.Fatalf("%s != %s", Axel.Animal(), s)
	}
}

func TestAxelName(t *testing.T) {
	if ok := Axel.Name() == axel; !ok {
		t.Fatalf("%s != %s", Axel.Name(), axel)
	}
}
