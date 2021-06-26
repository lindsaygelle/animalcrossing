package sheep

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCurlosAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Curlos.Animal() == s; !ok {
		t.Fatalf("%s != %s", Curlos.Animal(), s)
	}
}

func TestCurlosName(t *testing.T) {
	if ok := Curlos.Name() == curlos; !ok {
		t.Fatalf("%s != %s", Curlos.Name(), curlos)
	}
}
