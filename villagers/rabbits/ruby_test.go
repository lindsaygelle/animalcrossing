package rabbits

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRubyAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Ruby.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ruby.Animal(), s)
	}
}

func TestRubyName(t *testing.T) {
	if ok := Ruby.Name() == ruby; !ok {
		t.Fatalf("%s != %s", Ruby.Name(), ruby)
	}
}
