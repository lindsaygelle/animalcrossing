package koalas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGonzoAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Gonzo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gonzo.Animal(), s)
	}
}

func TestGonzoName(t *testing.T) {
	if ok := Gonzo.Name() == gonzo; !ok {
		t.Fatalf("%s != %s", Gonzo.Name(), gonzo)
	}
}
