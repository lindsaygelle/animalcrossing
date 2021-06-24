package goats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNanAnimal(t *testing.T) {
	var s string = animals.Goat.Name()
	if ok := Nan.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nan.Animal(), s)
	}
}

func TestNanName(t *testing.T) {
	if ok := Nan.Name() == nan; !ok {
		t.Fatalf("%s != %s", Nan.Name(), nan)
	}
}
