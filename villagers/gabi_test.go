package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGabiAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Gabi.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gabi.Animal(), s)
	}
}

func TestGabiName(t *testing.T) {
	if ok := Gabi.Name() == gabi; !ok {
		t.Fatalf("%s != %s", Gabi.Name(), gabi)
	}
}
