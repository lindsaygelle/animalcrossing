package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCamofrogAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Camofrog.Animal() == s; !ok {
		t.Fatalf("%s != %s", Camofrog.Animal(), s)
	}
}

func TestCamofrogName(t *testing.T) {
	if ok := Camofrog.Name() == camofrog; !ok {
		t.Fatalf("%s != %s", Camofrog.Name(), camofrog)
	}
}
