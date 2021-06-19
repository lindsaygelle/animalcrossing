package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRoscoeAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Roscoe.Animal() == s; !ok {
		t.Fatalf("%s != %s", Roscoe.Animal(), s)
	}
}

func TestRoscoeName(t *testing.T) {
	if ok := Roscoe.Name() == roscoe; !ok {
		t.Fatalf("%s != %s", Roscoe.Name(), roscoe)
	}
}
