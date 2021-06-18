package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKiddAnimal(t *testing.T) {
	var s string = animals.Goat.Name()
	if ok := Kidd.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kidd.Animal(), s)
	}
}

func TestKiddName(t *testing.T) {
	if ok := Kidd.Name() == kidd; !ok {
		t.Fatalf("%s != %s", Kidd.Name(), kidd)
	}
}
