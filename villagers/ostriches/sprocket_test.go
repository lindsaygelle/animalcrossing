package ostriches

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSprocketAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Sprocket.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sprocket.Animal(), s)
	}
}

func TestSprocketName(t *testing.T) {
	if ok := Sprocket.Name() == sprocket; !ok {
		t.Fatalf("%s != %s", Sprocket.Name(), sprocket)
	}
}
