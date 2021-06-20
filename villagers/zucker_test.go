package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestZuckerAnimal(t *testing.T) {
	var s string = animals.Octopus.Name()
	if ok := Zucker.Animal() == s; !ok {
		t.Fatalf("%s != %s", Zucker.Animal(), s)
	}
}

func TestZuckerName(t *testing.T) {
	if ok := Zucker.Name() == zucker; !ok {
		t.Fatalf("%s != %s", Zucker.Name(), zucker)
	}
}
