package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCubeAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Cube.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cube.Animal(), s)
	}
}

func TestCubeName(t *testing.T) {
	if ok := Cube.Name() == cube; !ok {
		t.Fatalf("%s != %s", Cube.Name(), cube)
	}
}
