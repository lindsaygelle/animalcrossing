package kappas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLeilaniAnimal(t *testing.T) {
	var s string = animals.Kappa.Name()
	if ok := Leilani.Animal() == s; !ok {
		t.Fatalf("%s != %s", Leilani.Animal(), s)
	}
}

func TestLeilaniName(t *testing.T) {
	if ok := Leilani.Name() == leilani; !ok {
		t.Fatalf("%s != %s", Leilani.Name(), leilani)
	}
}
