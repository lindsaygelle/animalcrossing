package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBellaAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Bella.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bella.Animal(), s)
	}
}

func TestBellaName(t *testing.T) {
	if ok := Bella.Name() == bella; !ok {
		t.Fatalf("%s != %s", Bella.Name(), bella)
	}
}
