package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOctavianAnimal(t *testing.T) {
	var s string = animals.Octopus.Name()
	if ok := Octavian.Animal() == s; !ok {
		t.Fatalf("%s != %s", Octavian.Animal(), s)
	}
}

func TestOctavianName(t *testing.T) {
	if ok := Octavian.Name() == octavian; !ok {
		t.Fatalf("%s != %s", Octavian.Name(), octavian)
	}
}
