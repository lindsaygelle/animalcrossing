package koalas

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAliceAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Alice.Animal() == s; !ok {
		t.Fatalf("%s != %s", Alice.Animal(), s)
	}
}

func TestAliceName(t *testing.T) {
	if ok := Alice.Name() == alice; !ok {
		t.Fatalf("%s != %s", Alice.Name(), alice)
	}
}
