package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPeggyAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Peggy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Peggy.Animal(), s)
	}
}

func TestPeggyName(t *testing.T) {
	if ok := Peggy.Name() == peggy; !ok {
		t.Fatalf("%s != %s", Peggy.Name(), peggy)
	}
}
