package rabbits

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestHopkinsAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Hopkins.Animal() == s; !ok {
		t.Fatalf("%s != %s", Hopkins.Animal(), s)
	}
}

func TestHopkinsName(t *testing.T) {
	if ok := Hopkins.Name() == hopkins; !ok {
		t.Fatalf("%s != %s", Hopkins.Name(), hopkins)
	}
}
