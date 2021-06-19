package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestClydeAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Clyde.Animal() == s; !ok {
		t.Fatalf("%s != %s", Clyde.Animal(), s)
	}
}

func TestClydeName(t *testing.T) {
	if ok := Clyde.Name() == clyde; !ok {
		t.Fatalf("%s != %s", Clyde.Name(), clyde)
	}
}
