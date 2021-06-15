package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRobinName(t *testing.T) {
	if ok := Robin.Name() == robin; !ok {
		t.Fatalf("%s != %s", Robin.Name(), robin)
	}
}

func TestRobinSpecies(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Robin.Animal() == s; !ok {
		t.Fatalf("%s != %s", Robin.Animal(), s)
	}
}
