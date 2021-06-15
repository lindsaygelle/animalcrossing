package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestReeseAnimal(t *testing.T) {
	var s string = animals.Alpaca.Name()
	if ok := Reese.Animal() == s; !ok {
		t.Fatalf("%s != %s", Reese.Animal(), s)
	}
}

func TestReeseName(t *testing.T) {
	if ok := Reese.Name() == reese; !ok {
		t.Fatalf("%s != %s", Reese.Name(), reese)
	}
}
