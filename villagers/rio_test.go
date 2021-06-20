package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRioAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Rio.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rio.Animal(), s)
	}
}

func TestRioName(t *testing.T) {
	if ok := Rio.Name() == rio; !ok {
		t.Fatalf("%s != %s", Rio.Name(), rio)
	}
}
