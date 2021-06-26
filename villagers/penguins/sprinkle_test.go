package penguins

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSprinkleAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Sprinkle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sprinkle.Animal(), s)
	}
}

func TestSprinkleName(t *testing.T) {
	if ok := Sprinkle.Name() == sprinkle; !ok {
		t.Fatalf("%s != %s", Sprinkle.Name(), sprinkle)
	}
}
