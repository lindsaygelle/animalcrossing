package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBigTopAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := BigTop.Animal() == s; !ok {
		t.Fatalf("%s != %s", BigTop.Animal(), s)
	}
}

func TestBigTopName(t *testing.T) {
	if ok := BigTop.Name() == bigTop; !ok {
		t.Fatalf("%s != %s", BigTop.Name(), bigTop)
	}
}
