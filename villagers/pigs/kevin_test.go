package pigs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestKevinAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Kevin.Animal() == s; !ok {
		t.Fatalf("%s != %s", Kevin.Animal(), s)
	}
}

func TestKevinName(t *testing.T) {
	if ok := Kevin.Name() == kevin; !ok {
		t.Fatalf("%s != %s", Kevin.Name(), kevin)
	}
}
