package gorillas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPeeweeAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Peewee.Animal() == s; !ok {
		t.Fatalf("%s != %s", Peewee.Animal(), s)
	}
}

func TestPeeweeName(t *testing.T) {
	if ok := Peewee.Name() == peewee; !ok {
		t.Fatalf("%s != %s", Peewee.Name(), peewee)
	}
}
