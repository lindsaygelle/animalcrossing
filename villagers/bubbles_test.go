package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBubblesAnimal(t *testing.T) {
	var s string = animals.Hippo.Name()
	if ok := Bubbles.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bubbles.Animal(), s)
	}
}

func TestBubblesName(t *testing.T) {
	if ok := Bubbles.Name() == bubbles; !ok {
		t.Fatalf("%s != %s", Bubbles.Name(), bubbles)
	}
}
