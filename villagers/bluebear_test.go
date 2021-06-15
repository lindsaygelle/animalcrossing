package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBluebearName(t *testing.T) {
	if ok := Bluebear.Name() == bluebear; !ok {
		t.Fatalf("%s != %s", Bluebear.Name(), bluebear)
	}
}

func TestBluebearSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Bluebear.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bluebear.Animal(), s)
	}
}
