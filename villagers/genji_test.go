package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGenjiAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Genji.Animal() == s; !ok {
		t.Fatalf("%s != %s", Genji.Animal(), s)
	}
}

func TestGenjiName(t *testing.T) {
	if ok := Genji.Name() == genji; !ok {
		t.Fatalf("%s != %s", Genji.Name(), genji)
	}
}
