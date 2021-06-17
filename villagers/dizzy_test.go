package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDizzyAnimal(t *testing.T) {
	var s string = animals.Elephant.Name()
	if ok := Dizzy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Dizzy.Animal(), s)
	}
}

func TestDizzyName(t *testing.T) {
	if ok := Dizzy.Name() == dizzy; !ok {
		t.Fatalf("%s != %s", Dizzy.Name(), dizzy)
	}
}
