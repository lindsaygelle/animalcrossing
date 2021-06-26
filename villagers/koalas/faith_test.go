package koalas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFaithAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Faith.Animal() == s; !ok {
		t.Fatalf("%s != %s", Faith.Animal(), s)
	}
}

func TestFaithName(t *testing.T) {
	if ok := Faith.Name() == faith; !ok {
		t.Fatalf("%s != %s", Faith.Name(), faith)
	}
}
