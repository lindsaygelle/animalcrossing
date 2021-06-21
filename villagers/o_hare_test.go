package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOHareAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := OHare.Animal() == s; !ok {
		t.Fatalf("%s != %s", OHare.Animal(), s)
	}
}

func TestOHareName(t *testing.T) {
	if ok := OHare.Name() == oHare; !ok {
		t.Fatalf("%s != %s", OHare.Name(), oHare)
	}
}
