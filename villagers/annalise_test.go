package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAnnaliseAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Annalise.Animal() == s; !ok {
		t.Fatalf("%s != %s", Annalise.Animal(), s)
	}
}

func TestAnnaliseName(t *testing.T) {
	if ok := Annalise.Name() == annalise; !ok {
		t.Fatalf("%s != %s", Annalise.Name(), annalise)
	}
}
