package cows

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPattyAnimal(t *testing.T) {
	var s string = animals.Cow.Name()
	if ok := Patty.Animal() == s; !ok {
		t.Fatalf("%s != %s", Patty.Animal(), s)
	}
}

func TestPattyName(t *testing.T) {
	if ok := Patty.Name() == patty; !ok {
		t.Fatalf("%s != %s", Patty.Name(), patty)
	}
}
