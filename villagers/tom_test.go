package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTomAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Tom.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tom.Animal(), s)
	}
}

func TestTomName(t *testing.T) {
	if ok := Tom.Name() == tom; !ok {
		t.Fatalf("%s != %s", Tom.Name(), tom)
	}
}
