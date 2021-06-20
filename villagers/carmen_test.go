package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCarmenAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Carmen.Animal() == s; !ok {
		t.Fatalf("%s != %s", Carmen.Animal(), s)
	}
}

func TestCarmenName(t *testing.T) {
	if ok := Carmen.Name() == carmen; !ok {
		t.Fatalf("%s != %s", Carmen.Name(), carmen)
	}
}
