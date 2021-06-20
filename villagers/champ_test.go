package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestChampAnimal(t *testing.T) {
	var s string = animals.Monkey.Name()
	if ok := Champ.Animal() == s; !ok {
		t.Fatalf("%s != %s", Champ.Animal(), s)
	}
}

func TestChampName(t *testing.T) {
	if ok := Champ.Name() == champ; !ok {
		t.Fatalf("%s != %s", Champ.Name(), champ)
	}
}
