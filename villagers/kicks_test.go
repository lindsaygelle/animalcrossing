package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKicksAnimal(t *testing.T) {
	var s string = animals.Skunk.Name()
	if ok := Kicks.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kicks.Animal(), s)
	}
}

func TestKicksName(t *testing.T) {
	if ok := Kicks.Name() == kicks; !ok {
		t.Fatalf("%s != %s", Kicks.Name(), kicks)
	}
}
