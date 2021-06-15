package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKodyName(t *testing.T) {
	if ok := Kody.Name() == kody; !ok {
		t.Fatalf("%s != %s", Kody.Name(), kody)
	}
}

func TestKodySpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Kody.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kody.Animal(), s)
	}
}
