package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAnkhaAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Ankha.Animal() == s; !ok {
		t.Fatalf("%s != %s", Ankha.Animal(), s)
	}
}

func TestAnkhaName(t *testing.T) {
	if ok := Ankha.Name() == ankha; !ok {
		t.Fatalf("%s != %s", Ankha.Name(), ankha)
	}
}
