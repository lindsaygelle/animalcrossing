package anteaters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLuluAnimal(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Lulu.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lulu.Animal(), s)
	}
}

func TestLuluName(t *testing.T) {
	if ok := Lulu.Name() == lulu; !ok {
		t.Fatalf("%s != %s", Lulu.Name(), lulu)
	}
}
