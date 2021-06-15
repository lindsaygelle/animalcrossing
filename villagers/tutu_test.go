package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTutuName(t *testing.T) {
	if ok := Tutu.Name() == tutu; !ok {
		t.Fatalf("%s != %s", Tutu.Name(), tutu)
	}
}

func TestTutuSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Tutu.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tutu.Animal(), s)
	}
}
