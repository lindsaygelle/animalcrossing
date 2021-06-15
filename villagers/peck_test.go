package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPeckAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Peck.Animal() == s; !ok {
		t.Fatalf("%s != %s", Peck.Animal(), s)
	}
}

func TestPeckName(t *testing.T) {
	if ok := Peck.Name() == peck; !ok {
		t.Fatalf("%s != %s", Peck.Name(), peck)
	}
}
