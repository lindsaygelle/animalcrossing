package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHazelAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Hazel.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hazel.Animal(), s)
	}
}

func TestHazelName(t *testing.T) {
	if ok := Hazel.Name() == hazel; !ok {
		t.Fatalf("%s != %s", Hazel.Name(), hazel)
	}
}
