package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFrigaAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Friga.Animal() == s; !ok {
		t.Fatalf("%s != %s", Friga.Animal(), s)
	}
}

func TestFrigaName(t *testing.T) {
	if ok := Friga.Name() == friga; !ok {
		t.Fatalf("%s != %s", Friga.Name(), friga)
	}
}
