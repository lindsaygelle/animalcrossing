package goats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPashminaAnimal(t *testing.T) {
	var s string = animals.Goat.Name()
	if ok := Pashmina.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pashmina.Animal(), s)
	}
}

func TestPashminaName(t *testing.T) {
	if ok := Pashmina.Name() == pashmina; !ok {
		t.Fatalf("%s != %s", Pashmina.Name(), pashmina)
	}
}
