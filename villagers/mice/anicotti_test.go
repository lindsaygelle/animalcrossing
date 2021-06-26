package mice

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAnicottiAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Anicotti.Animal() == s; !ok {
		t.Fatalf("%s != %s", Anicotti.Animal(), s)
	}
}

func TestAnicottiName(t *testing.T) {
	if ok := Anicotti.Name() == anicotti; !ok {
		t.Fatalf("%s != %s", Anicotti.Name(), anicotti)
	}
}
