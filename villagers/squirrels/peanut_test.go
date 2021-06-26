package squirrels

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPeanutAnimal(t *testing.T) {
	var s string = animals.Squirrel.Name()
	if ok := Peanut.Animal() == s; !ok {
		t.Fatalf("%s != %s", Peanut.Animal(), s)
	}
}

func TestPeanutName(t *testing.T) {
	if ok := Peanut.Name() == peanut; !ok {
		t.Fatalf("%s != %s", Peanut.Name(), peanut)
	}
}
