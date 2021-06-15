package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChesterAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Chester.Animal() == s; !ok {
		t.Fatalf("%s != %s", Chester.Animal(), s)
	}
}

func TestChesterName(t *testing.T) {
	if ok := Chester.Name() == chester; !ok {
		t.Fatalf("%s != %s", Chester.Name(), chester)
	}
}
