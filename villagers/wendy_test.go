package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWendyAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Wendy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Wendy.Animal(), s)
	}
}

func TestWendyName(t *testing.T) {
	if ok := Wendy.Name() == wendy; !ok {
		t.Fatalf("%s != %s", Wendy.Name(), wendy)
	}
}
