package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLilyAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Lily.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lily.Animal(), s)
	}
}

func TestLilyName(t *testing.T) {
	if ok := Lily.Name() == lily; !ok {
		t.Fatalf("%s != %s", Lily.Name(), lily)
	}
}
