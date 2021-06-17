package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPrinceAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Prince.Animal() == s; !ok {
		t.Fatalf("%s != %s", Prince.Animal(), s)
	}
}

func TestPrinceName(t *testing.T) {
	if ok := Prince.Name() == prince; !ok {
		t.Fatalf("%s != %s", Prince.Name(), prince)
	}
}
