package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLuckyAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Lucky.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lucky.Animal(), s)
	}
}

func TestLuckyName(t *testing.T) {
	if ok := Lucky.Name() == lucky; !ok {
		t.Fatalf("%s != %s", Lucky.Name(), lucky)
	}
}
