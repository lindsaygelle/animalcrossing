package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHamletAnimal(t *testing.T) {
	var s string = animals.Hamster.Name()
	if ok := Hamlet.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hamlet.Animal(), s)
	}
}

func TestHamletName(t *testing.T) {
	if ok := Hamlet.Name() == hamlet; !ok {
		t.Fatalf("%s != %s", Hamlet.Name(), hamlet)
	}
}
