package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBangleAnimal(t *testing.T) {
	var s string = animals.Tiger.Name()
	if ok := Bangle.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bangle.Animal(), s)
	}
}

func TestBangleName(t *testing.T) {
	if ok := Bangle.Name() == bangle; !ok {
		t.Fatalf("%s != %s", Bangle.Name(), bangle)
	}
}
