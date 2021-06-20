package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKappnAnimal(t *testing.T) {
	var s string = animals.Kappa.Name()
	if ok := Kappn.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kappn.Animal(), s)
	}
}

func TestKappnName(t *testing.T) {
	if ok := Kappn.Name() == kappn; !ok {
		t.Fatalf("%s != %s", Kappn.Name(), kappn)
	}
}
