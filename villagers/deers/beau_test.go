package deers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBeauAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Beau.Animal() == s; !ok {
		t.Fatalf("%s != %s", Beau.Animal(), s)
	}
}

func TestBeauName(t *testing.T) {
	if ok := Beau.Name() == beau; !ok {
		t.Fatalf("%s != %s", Beau.Name(), beau)
	}
}
