package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTrufflesAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Truffles.Animal() == s; !ok {
		t.Fatalf("%s != %s", Truffles.Animal(), s)
	}
}

func TestTrufflesName(t *testing.T) {
	if ok := Truffles.Name() == truffles; !ok {
		t.Fatalf("%s != %s", Truffles.Name(), truffles)
	}
}
