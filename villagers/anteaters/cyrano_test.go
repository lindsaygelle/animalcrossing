package anteaters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCyranoAnimal(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Cyrano.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cyrano.Animal(), s)
	}
}

func TestCyranoName(t *testing.T) {
	if ok := Cyrano.Name() == cyrano; !ok {
		t.Fatalf("%s != %s", Cyrano.Name(), cyrano)
	}
}
