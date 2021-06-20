package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPuckAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Puck.Animal() == s; !ok {
		t.Fatalf("%s != %s", Puck.Animal(), s)
	}
}

func TestPuckName(t *testing.T) {
	if ok := Puck.Name() == puck; !ok {
		t.Fatalf("%s != %s", Puck.Name(), puck)
	}
}
