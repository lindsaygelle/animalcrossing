package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChowName(t *testing.T) {
	if ok := Chow.Name() == chow; !ok {
		t.Fatalf("%s != %s", Chow.Name(), chow)
	}
}

func TestChowSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Chow.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chow.Animal(), s)
	}
}
