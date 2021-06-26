package penguins

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBoomerAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Boomer.Animal() == s; !ok {
		t.Fatalf("%s != %s", Boomer.Animal(), s)
	}
}

func TestBoomerName(t *testing.T) {
	if ok := Boomer.Name() == boomer; !ok {
		t.Fatalf("%s != %s", Boomer.Name(), boomer)
	}
}
