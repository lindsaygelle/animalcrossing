package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBeppeName(t *testing.T) {
	if ok := Beppe.Name() == beppe; !ok {
		t.Fatalf("%s != %s", Beppe.Name(), beppe)
	}
}

func TestBeppeSpecies(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Beppe.Animal() == s; !ok {
		t.Fatalf("%s != %s", Beppe.Animal(), s)
	}
}
