package monkeys

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNanaAnimal(t *testing.T) {
	var s string = animals.Monkey.Name()
	if ok := Nana.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nana.Animal(), s)
	}
}

func TestNanaName(t *testing.T) {
	if ok := Nana.Name() == nana; !ok {
		t.Fatalf("%s != %s", Nana.Name(), nana)
	}
}
