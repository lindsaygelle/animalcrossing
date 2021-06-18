package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAppleAnimal(t *testing.T) {
	var s string = animals.Hamster.Name()
	if ok := Apple.Animal() == s; !ok {
		t.Fatalf("%s != %s", Apple.Animal(), s)
	}
}

func TestAppleName(t *testing.T) {
	if ok := Apple.Name() == apple; !ok {
		t.Fatalf("%s != %s", Apple.Name(), apple)
	}
}
