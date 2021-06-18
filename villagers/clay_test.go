package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestClayAnimal(t *testing.T) {
	var s string = animals.Hamster.Name()
	if ok := Clay.Animal() == s; !ok {
		t.Fatalf("%s != %s", Clay.Animal(), s)
	}
}

func TestClayName(t *testing.T) {
	if ok := Clay.Name() == clay; !ok {
		t.Fatalf("%s != %s", Clay.Name(), clay)
	}
}
