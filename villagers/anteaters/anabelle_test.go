package anteaters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAnabelleAnimal(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Anabelle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Anabelle.Animal(), s)
	}
}

func TestAnabelleName(t *testing.T) {
	if ok := Anabelle.Name() == anabelle; !ok {
		t.Fatalf("%s != %s", Anabelle.Name(), anabelle)
	}
}
