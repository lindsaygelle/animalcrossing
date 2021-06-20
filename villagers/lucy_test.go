package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLucyAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Lucy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lucy.Animal(), s)
	}
}

func TestLucyName(t *testing.T) {
	if ok := Lucy.Name() == lucy; !ok {
		t.Fatalf("%s != %s", Lucy.Name(), lucy)
	}
}
