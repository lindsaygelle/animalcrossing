package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTipperAnimal(t *testing.T) {
	var s string = animals.Cow.Name()
	if ok := Tipper.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tipper.Animal(), s)
	}
}

func TestTipperName(t *testing.T) {
	if ok := Tipper.Name() == tipper; !ok {
		t.Fatalf("%s != %s", Tipper.Name(), tipper)
	}
}
