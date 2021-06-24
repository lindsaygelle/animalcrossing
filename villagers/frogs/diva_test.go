package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDivaAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Diva.Animal() == s; !ok {
		t.Fatalf("%s != %s", Diva.Animal(), s)
	}
}

func TestDivaName(t *testing.T) {
	if ok := Diva.Name() == diva; !ok {
		t.Fatalf("%s != %s", Diva.Name(), diva)
	}
}
