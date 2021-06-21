package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGulliverAnimal(t *testing.T) {
	var s string = animals.SeaGull.Name()
	if ok := Gulliver.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gulliver.Animal(), s)
	}
}

func TestGulliverName(t *testing.T) {
	if ok := Gulliver.Name() == gulliver; !ok {
		t.Fatalf("%s != %s", Gulliver.Name(), gulliver)
	}
}
