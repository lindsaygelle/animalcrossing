package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDobieAnimal(t *testing.T) {
	var s string = animals.Wolf.Name()
	if ok := Dobie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Dobie.Animal(), s)
	}
}

func TestDobieName(t *testing.T) {
	if ok := Dobie.Name() == dobie; !ok {
		t.Fatalf("%s != %s", Dobie.Name(), dobie)
	}
}
