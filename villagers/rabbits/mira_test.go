package rabbits

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMiraAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Mira.Animal() == s; !ok {
		t.Fatalf("%s != %s", Mira.Animal(), s)
	}
}

func TestMiraName(t *testing.T) {
	if ok := Mira.Name() == mira; !ok {
		t.Fatalf("%s != %s", Mira.Name(), mira)
	}
}
