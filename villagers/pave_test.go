package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPaveAnimal(t *testing.T) {
	var s string = animals.Peacock.Name()
	if ok := Pave.Animal() == s; !ok {
		t.Fatalf("%s != %s", Pave.Animal(), s)
	}
}

func TestPaveName(t *testing.T) {
	if ok := Pave.Name() == pave; !ok {
		t.Fatalf("%s != %s", Pave.Name(), pave)
	}
}
