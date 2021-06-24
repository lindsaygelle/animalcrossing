package goats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBillyAnimal(t *testing.T) {
	var s string = animals.Goat.Name()
	if ok := Billy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Billy.Animal(), s)
	}
}

func TestBillyName(t *testing.T) {
	if ok := Billy.Name() == billy; !ok {
		t.Fatalf("%s != %s", Billy.Name(), billy)
	}
}
