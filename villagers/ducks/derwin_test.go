package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDerwinAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Derwin.Animal() == s; !ok {
		t.Fatalf("%s != %s", Derwin.Animal(), s)
	}
}

func TestDerwinName(t *testing.T) {
	if ok := Derwin.Name() == derwin; !ok {
		t.Fatalf("%s != %s", Derwin.Name(), derwin)
	}
}
