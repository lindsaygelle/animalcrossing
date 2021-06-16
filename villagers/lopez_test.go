package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLopezAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Lopez.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lopez.Animal(), s)
	}
}

func TestLopezName(t *testing.T) {
	if ok := Lopez.Name() == lopez; !ok {
		t.Fatalf("%s != %s", Lopez.Name(), lopez)
	}
}
