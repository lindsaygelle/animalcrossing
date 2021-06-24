package birds

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJoeAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Joe.Animal() == s; !ok {
		t.Fatalf("%s != %s", Joe.Animal(), s)
	}
}

func TestJoeName(t *testing.T) {
	if ok := Joe.Name() == joe; !ok {
		t.Fatalf("%s != %s", Joe.Name(), joe)
	}
}
