package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLeighAnimal(t *testing.T) {
	var s string = animals.Chicken.Name()
	if ok := Leigh.Animal() == s; !ok {
		t.Fatalf("%s != %s", Leigh.Animal(), s)
	}
}

func TestLeighName(t *testing.T) {
	if ok := Leigh.Name() == leigh; !ok {
		t.Fatalf("%s != %s", Leigh.Name(), leigh)
	}
}
