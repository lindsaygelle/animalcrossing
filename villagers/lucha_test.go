package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLuchaName(t *testing.T) {
	if ok := Lucha.Name() == lucha; !ok {
		t.Fatalf("%s != %s", Lucha.Name(), lucha)
	}
}

func TestLuchaSpecies(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Lucha.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lucha.Animal(), s)
	}
}
