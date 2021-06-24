package deers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestZellAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Zell.Animal() == s; !ok {
		t.Fatalf("%s != %s", Zell.Animal(), s)
	}
}

func TestZellName(t *testing.T) {
	if ok := Zell.Name() == zell; !ok {
		t.Fatalf("%s != %s", Zell.Name(), zell)
	}
}
