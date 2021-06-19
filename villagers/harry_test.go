package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHarryAnimal(t *testing.T) {
	var s string = animals.Hippo.Name()
	if ok := Harry.Animal() == s; !ok {
		t.Fatalf("%s != %s", Harry.Animal(), s)
	}
}

func TestHarryName(t *testing.T) {
	if ok := Harry.Name() == harry; !ok {
		t.Fatalf("%s != %s", Harry.Name(), harry)
	}
}
