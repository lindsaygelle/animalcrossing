package koalas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMelbaAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Melba.Animal() == s; !ok {
		t.Fatalf("%s != %s", Melba.Animal(), s)
	}
}

func TestMelbaName(t *testing.T) {
	if ok := Melba.Name() == melba; !ok {
		t.Fatalf("%s != %s", Melba.Name(), melba)
	}
}
