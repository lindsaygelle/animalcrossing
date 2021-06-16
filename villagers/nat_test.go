package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNatAnimal(t *testing.T) {
	var s string = animals.Chameleon.Name()
	if ok := Nat.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nat.Animal(), s)
	}
}

func TestNatName(t *testing.T) {
	if ok := Nat.Name() == nat; !ok {
		t.Fatalf("%s != %s", Nat.Name(), nat)
	}
}
