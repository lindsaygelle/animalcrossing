package koalas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOzzieAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Ozzie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ozzie.Animal(), s)
	}
}

func TestOzzieName(t *testing.T) {
	if ok := Ozzie.Name() == ozzie; !ok {
		t.Fatalf("%s != %s", Ozzie.Name(), ozzie)
	}
}
