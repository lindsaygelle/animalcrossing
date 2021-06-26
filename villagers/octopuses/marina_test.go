package octopuses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMarinaAnimal(t *testing.T) {
	var s string = animals.Octopus.Name()
	if ok := Marina.Animal() == s; !ok {
		t.Fatalf("%s != %s", Marina.Animal(), s)
	}
}

func TestMarinaName(t *testing.T) {
	if ok := Marina.Name() == marina; !ok {
		t.Fatalf("%s != %s", Marina.Name(), marina)
	}
}
