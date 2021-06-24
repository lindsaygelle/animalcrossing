package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestQuillsonAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Quillson.Animal() == s; !ok {
		t.Fatalf("%s != %s", Quillson.Animal(), s)
	}
}

func TestQuillsonName(t *testing.T) {
	if ok := Quillson.Name() == quillson; !ok {
		t.Fatalf("%s != %s", Quillson.Name(), quillson)
	}
}
