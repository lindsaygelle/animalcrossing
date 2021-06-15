package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMurphyAnimal(t *testing.T) {
	var s string = animals.Bear.Name()
	if ok := Murphy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Murphy.Animal(), s)
	}
}

func TestMurphyName(t *testing.T) {
	if ok := Murphy.Name() == murphy; !ok {
		t.Fatalf("%s != %s", Murphy.Name(), murphy)
	}
}
