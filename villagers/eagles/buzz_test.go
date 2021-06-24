package eagles

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBuzzAnimal(t *testing.T) {
	var s string = animals.Eagle.Name()
	if ok := Buzz.Animal() == s; !ok {
		t.Fatalf("%s != %s", Buzz.Animal(), s)
	}
}

func TestBuzzName(t *testing.T) {
	if ok := Buzz.Name() == buzz; !ok {
		t.Fatalf("%s != %s", Buzz.Name(), buzz)
	}
}
