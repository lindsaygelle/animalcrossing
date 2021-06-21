package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPietroAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Pietro.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pietro.Animal(), s)
	}
}

func TestPietroName(t *testing.T) {
	if ok := Pietro.Name() == pietro; !ok {
		t.Fatalf("%s != %s", Pietro.Name(), pietro)
	}
}
