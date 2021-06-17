package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWeberAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Weber.Animal() == s; !ok {
		t.Fatalf("%s != %s", Weber.Animal(), s)
	}
}

func TestWeberName(t *testing.T) {
	if ok := Weber.Name() == weber; !ok {
		t.Fatalf("%s != %s", Weber.Name(), weber)
	}
}
