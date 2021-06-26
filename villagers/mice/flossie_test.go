package mice

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFlossieAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Flossie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Flossie.Animal(), s)
	}
}

func TestFlossieName(t *testing.T) {
	if ok := Flossie.Name() == flossie; !ok {
		t.Fatalf("%s != %s", Flossie.Name(), flossie)
	}
}
