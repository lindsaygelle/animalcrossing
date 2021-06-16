package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDeirdreAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Deirdre.Animal() == s; !ok {
		t.Fatalf("%s != %s", Deirdre.Animal(), s)
	}
}

func TestDeirdreName(t *testing.T) {
	if ok := Deirdre.Name() == deirdre; !ok {
		t.Fatalf("%s != %s", Deirdre.Name(), deirdre)
	}
}
