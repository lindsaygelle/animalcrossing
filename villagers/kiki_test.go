package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKikiAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Kiki.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kiki.Animal(), s)
	}
}

func TestKikiName(t *testing.T) {
	if ok := Kiki.Name() == kiki; !ok {
		t.Fatalf("%s != %s", Kiki.Name(), kiki)
	}
}
