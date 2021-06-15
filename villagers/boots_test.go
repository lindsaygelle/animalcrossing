package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBootsName(t *testing.T) {
	if ok := Boots.Name() == boots; !ok {
		t.Fatalf("%s != %s", Boots.Name(), boots)
	}
}

func TestBootsSpecies(t *testing.T) {
	var s string = animals.Alligator.Name()
	if ok := Boots.Animal() == s; !ok {
		t.Fatalf("%s != %s", Boots.Animal(), s)
	}
}
