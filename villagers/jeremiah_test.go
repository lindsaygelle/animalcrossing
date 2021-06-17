package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJeremiahAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Jeremiah.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jeremiah.Animal(), s)
	}
}

func TestJeremiahName(t *testing.T) {
	if ok := Jeremiah.Name() == jeremiah; !ok {
		t.Fatalf("%s != %s", Jeremiah.Name(), jeremiah)
	}
}
