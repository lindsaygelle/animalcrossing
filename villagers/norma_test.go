package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNormaAnimal(t *testing.T) {
	var s string = animals.Cow.Name()
	if ok := Norma.Animal() == s; !ok {
		t.Fatalf("%s != %s", Norma.Animal(), s)
	}
}

func TestNormaName(t *testing.T) {
	if ok := Norma.Name() == norma; !ok {
		t.Fatalf("%s != %s", Norma.Name(), norma)
	}
}
