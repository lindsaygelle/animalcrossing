package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPhineasAnimal(t *testing.T) {
	var s string = animals.FurSeal.Name()
	if ok := Phineas.Animal() == s; !ok {
		t.Fatalf("%s != %s", Phineas.Animal(), s)
	}
}

func TestPhineasName(t *testing.T) {
	if ok := Phineas.Name() == phineas; !ok {
		t.Fatalf("%s != %s", Phineas.Name(), phineas)
	}
}
