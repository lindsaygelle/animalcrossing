package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestNindoriAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Nindori.Animal() == s; !ok {
		t.Fatalf("%s != %s", Nindori.Animal(), s)
	}
}

func TestNindoriName(t *testing.T) {
	if ok := Nindori.Name() == nindori; !ok {
		t.Fatalf("%s != %s", Nindori.Name(), nindori)
	}
}
