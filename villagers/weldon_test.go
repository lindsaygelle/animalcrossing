package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestWeldonAnimal(t *testing.T) {
	var s string = animals.Bull.Name()
	if ok := Weldon.Animal() == s; !ok {
		t.Fatalf("%s != %s", Weldon.Animal(), s)
	}
}

func TestWeldonName(t *testing.T) {
	if ok := Weldon.Name() == weldon; !ok {
		t.Fatalf("%s != %s", Weldon.Name(), weldon)
	}
}
