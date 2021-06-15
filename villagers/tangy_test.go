package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTangyAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Tangy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tangy.Animal(), s)
	}
}

func TestTangyName(t *testing.T) {
	if ok := Tangy.Name() == tangy; !ok {
		t.Fatalf("%s != %s", Tangy.Name(), tangy)
	}
}
