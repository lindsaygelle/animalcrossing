package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJuliaAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Julia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Julia.Animal(), s)
	}
}

func TestJuliaName(t *testing.T) {
	if ok := Julia.Name() == julia; !ok {
		t.Fatalf("%s != %s", Julia.Name(), julia)
	}
}
