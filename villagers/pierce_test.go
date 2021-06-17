package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPierceAnimal(t *testing.T) {
	var s string = animals.Eagle.Name()
	if ok := Pierce.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pierce.Animal(), s)
	}
}

func TestPierceName(t *testing.T) {
	if ok := Pierce.Name() == pierce; !ok {
		t.Fatalf("%s != %s", Pierce.Name(), pierce)
	}
}
