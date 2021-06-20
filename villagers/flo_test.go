package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFloAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Flo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Flo.Animal(), s)
	}
}

func TestFloName(t *testing.T) {
	if ok := Flo.Name() == flo; !ok {
		t.Fatalf("%s != %s", Flo.Name(), flo)
	}
}
