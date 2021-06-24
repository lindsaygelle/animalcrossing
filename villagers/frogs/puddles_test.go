package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPuddlesAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Puddles.Animal() == s; !ok {
		t.Fatalf("%s != %s", Puddles.Animal(), s)
	}
}

func TestPuddlesName(t *testing.T) {
	if ok := Puddles.Name() == puddles; !ok {
		t.Fatalf("%s != %s", Puddles.Name(), puddles)
	}
}
