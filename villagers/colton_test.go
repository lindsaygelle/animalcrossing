package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestColtonAnimal(t *testing.T) {
	var s string = animals.Horse.Name()
	if ok := Colton.Animal() == s; !ok {
		t.Fatalf("%s != %s", Colton.Animal(), s)
	}
}

func TestColtonName(t *testing.T) {
	if ok := Colton.Name() == colton; !ok {
		t.Fatalf("%s != %s", Colton.Name(), colton)
	}
}
