package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJackAnimal(t *testing.T) {
	var s string = animals.Pumpkin.Name()
	if ok := Jack.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jack.Animal(), s)
	}
}

func TestJackName(t *testing.T) {
	if ok := Jack.Name() == jack; !ok {
		t.Fatalf("%s != %s", Jack.Name(), jack)
	}
}
