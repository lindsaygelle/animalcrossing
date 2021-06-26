package monkeys

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPorterAnimal(t *testing.T) {
	var s string = animals.Monkey.Name()
	if ok := Porter.Animal() == s; !ok {
		t.Fatalf("%s != %s", Porter.Animal(), s)
	}
}

func TestPorterName(t *testing.T) {
	if ok := Porter.Name() == porter; !ok {
		t.Fatalf("%s != %s", Porter.Name(), porter)
	}
}
