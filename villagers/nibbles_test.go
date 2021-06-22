package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNibblesAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Nibbles.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nibbles.Animal(), s)
	}
}

func TestNibblesName(t *testing.T) {
	if ok := Nibbles.Name() == nibbles; !ok {
		t.Fatalf("%s != %s", Nibbles.Name(), nibbles)
	}
}
