package dodos

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWilburAnimal(t *testing.T) {
	var s string = animals.Dodo.Name()
	if ok := Wilbur.Animal() == s; !ok {
		t.Fatalf("%s != %s", Wilbur.Animal(), s)
	}
}

func TestWilburName(t *testing.T) {
	if ok := Wilbur.Name() == wilbur; !ok {
		t.Fatalf("%s != %s", Wilbur.Name(), wilbur)
	}
}
