package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLymanAnimal(t *testing.T) {
	var s string = animals.Koala.Name()
	if ok := Lyman.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lyman.Animal(), s)
	}
}

func TestLymanName(t *testing.T) {
	if ok := Lyman.Name() == lyman; !ok {
		t.Fatalf("%s != %s", Lyman.Name(), lyman)
	}
}
