package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPascalAnimal(t *testing.T) {
	var s string = animals.Otter.Name()
	if ok := Pascal.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pascal.Animal(), s)
	}
}

func TestPascalName(t *testing.T) {
	if ok := Pascal.Name() == pascal; !ok {
		t.Fatalf("%s != %s", Pascal.Name(), pascal)
	}
}
