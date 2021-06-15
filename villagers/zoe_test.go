package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestZoeAnimal(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Zoe.Animal() == s; !ok {
		t.Fatalf("%s != %s", Zoe.Animal(), s)
	}
}

func TestZoeName(t *testing.T) {
	if ok := Zoe.Name() == zoe; !ok {
		t.Fatalf("%s != %s", Zoe.Name(), zoe)
	}
}
