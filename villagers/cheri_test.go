package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCheriName(t *testing.T) {
	if ok := Cheri.Name() == cheri; !ok {
		t.Fatalf("%s != %s", Cheri.Name(), cheri)
	}
}

func TestCheriSpecies(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Cheri.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cheri.Animal(), s)
	}
}
