package kappas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLeilaAnimal(t *testing.T) {
	var s string = animals.Kappa.Name()
	if ok := Leila.Animal() == s; !ok {
		t.Fatalf("%s != %s", Leila.Animal(), s)
	}
}

func TestLeilaName(t *testing.T) {
	if ok := Leila.Name() == leila; !ok {
		t.Fatalf("%s != %s", Leila.Name(), leila)
	}
}
