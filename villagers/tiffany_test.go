package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTiffanyAnimal(t *testing.T) {
	var s string = animals.Rabbit.Name()
	if ok := Tiffany.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tiffany.Animal(), s)
	}
}

func TestTiffanyName(t *testing.T) {
	if ok := Tiffany.Name() == tiffany; !ok {
		t.Fatalf("%s != %s", Tiffany.Name(), tiffany)
	}
}
