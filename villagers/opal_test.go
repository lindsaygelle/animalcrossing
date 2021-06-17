package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOpalAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Opal.Animal() == s; !ok {
		t.Fatalf("%s != %s", Opal.Animal(), s)
	}
}

func TestOpalName(t *testing.T) {
	if ok := Opal.Name() == opal; !ok {
		t.Fatalf("%s != %s", Opal.Name(), opal)
	}
}
