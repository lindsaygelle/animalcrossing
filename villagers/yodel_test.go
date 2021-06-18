package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestYodelAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Yodel.Animal() == s; !ok {
		t.Fatalf("%s != %s", Yodel.Animal(), s)
	}
}

func TestYodelName(t *testing.T) {
	if ok := Yodel.Name() == yodel; !ok {
		t.Fatalf("%s != %s", Yodel.Name(), yodel)
	}
}
