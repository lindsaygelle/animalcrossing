package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRudyAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Rudy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rudy.Animal(), s)
	}
}

func TestRudyName(t *testing.T) {
	if ok := Rudy.Name() == rudy; !ok {
		t.Fatalf("%s != %s", Rudy.Name(), rudy)
	}
}
