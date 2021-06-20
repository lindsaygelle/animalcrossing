package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTexAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Tex.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tex.Animal(), s)
	}
}

func TestTexName(t *testing.T) {
	if ok := Tex.Name() == tex; !ok {
		t.Fatalf("%s != %s", Tex.Name(), tex)
	}
}
