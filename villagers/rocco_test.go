package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRoccoAnimal(t *testing.T) {
	var s string = animals.Hippo.Name()
	if ok := Rocco.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rocco.Animal(), s)
	}
}

func TestRoccoName(t *testing.T) {
	if ok := Rocco.Name() == rocco; !ok {
		t.Fatalf("%s != %s", Rocco.Name(), rocco)
	}
}
