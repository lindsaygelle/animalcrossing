package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKatieAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Katie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Katie.Animal(), s)
	}
}

func TestKatieName(t *testing.T) {
	if ok := Katie.Name() == katie; !ok {
		t.Fatalf("%s != %s", Katie.Name(), katie)
	}
}
