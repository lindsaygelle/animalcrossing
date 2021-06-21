package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBonbonAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Bonbon.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bonbon.Animal(), s)
	}
}

func TestBonbonName(t *testing.T) {
	if ok := Bonbon.Name() == bonbon; !ok {
		t.Fatalf("%s != %s", Bonbon.Name(), bonbon)
	}
}
