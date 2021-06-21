package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEtoileAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Etoile.Animal() == s; !ok {
		t.Fatalf("%s != %s", Etoile.Animal(), s)
	}
}

func TestEtoileName(t *testing.T) {
	if ok := Etoile.Name() == etoile; !ok {
		t.Fatalf("%s != %s", Etoile.Name(), etoile)
	}
}
