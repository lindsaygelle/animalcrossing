package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRasherAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Rasher.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rasher.Animal(), s)
	}
}

func TestRasherName(t *testing.T) {
	if ok := Rasher.Name() == rasher; !ok {
		t.Fatalf("%s != %s", Rasher.Name(), rasher)
	}
}
