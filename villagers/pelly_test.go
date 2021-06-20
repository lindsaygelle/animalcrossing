package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPellyAnimal(t *testing.T) {
	var s string = animals.Pelican.Name()
	if ok := Pelly.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pelly.Animal(), s)
	}
}

func TestPellyName(t *testing.T) {
	if ok := Pelly.Name() == pelly; !ok {
		t.Fatalf("%s != %s", Pelly.Name(), pelly)
	}
}
