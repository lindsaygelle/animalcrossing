package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGastonAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Gaston.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gaston.Animal(), s)
	}
}

func TestGastonName(t *testing.T) {
	if ok := Gaston.Name() == gaston; !ok {
		t.Fatalf("%s != %s", Gaston.Name(), gaston)
	}
}
