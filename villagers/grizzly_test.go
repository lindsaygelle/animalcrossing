package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGrizzlyAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Grizzly.Animal() == s; !ok {
		t.Fatalf("%s != %s", Grizzly.Animal(), s)
	}
}

func TestGrizzlyName(t *testing.T) {
	if ok := Grizzly.Name() == grizzly; !ok {
		t.Fatalf("%s != %s", Grizzly.Name(), grizzly)
	}
}
