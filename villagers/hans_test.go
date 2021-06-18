package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHansAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Hans.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hans.Animal(), s)
	}
}

func TestHansName(t *testing.T) {
	if ok := Hans.Name() == hans; !ok {
		t.Fatalf("%s != %s", Hans.Name(), hans)
	}
}
