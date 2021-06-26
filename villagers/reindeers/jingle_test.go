package reindeers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJingleAnimal(t *testing.T) {
	var s string = animals.Reindeer.Name()
	if ok := Jingle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jingle.Animal(), s)
	}
}

func TestJingleName(t *testing.T) {
	if ok := Jingle.Name() == jingle; !ok {
		t.Fatalf("%s != %s", Jingle.Name(), jingle)
	}
}
