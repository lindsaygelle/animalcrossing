package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFaunaAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Fauna.Animal() == s; !ok {
		t.Fatalf("%s != %s", Fauna.Animal(), s)
	}
}

func TestFaunaName(t *testing.T) {
	if ok := Fauna.Name() == fauna; !ok {
		t.Fatalf("%s != %s", Fauna.Name(), fauna)
	}
}
