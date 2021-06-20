package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSimonAnimal(t *testing.T) {
	var s string = animals.Monkey.Name()
	if ok := Simon.Animal() == s; !ok {
		t.Fatalf("%s != %s", Simon.Animal(), s)
	}
}

func TestSimonName(t *testing.T) {
	if ok := Simon.Name() == simon; !ok {
		t.Fatalf("%s != %s", Simon.Name(), simon)
	}
}
