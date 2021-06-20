package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDoraAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Dora.Animal() == s; !ok {
		t.Fatalf("%s != %s", Dora.Animal(), s)
	}
}

func TestDoraName(t *testing.T) {
	if ok := Dora.Name() == dora; !ok {
		t.Fatalf("%s != %s", Dora.Name(), dora)
	}
}
