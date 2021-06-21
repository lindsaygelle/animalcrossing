package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMuffyAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Muffy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Muffy.Animal(), s)
	}
}

func TestMuffyName(t *testing.T) {
	if ok := Muffy.Name() == muffy; !ok {
		t.Fatalf("%s != %s", Muffy.Name(), muffy)
	}
}
