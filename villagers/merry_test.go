package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMerryAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Merry.Animal() == s; !ok {
		t.Fatalf("%s != %s", Merry.Animal(), s)
	}
}

func TestMerryName(t *testing.T) {
	if ok := Merry.Name() == merry; !ok {
		t.Fatalf("%s != %s", Merry.Name(), merry)
	}
}
