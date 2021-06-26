package pigs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestAgnesAnimal(t *testing.T) {
	var s string = animals.Pig.Name()
	if ok := Agnes.Animal() == s; !ok {
		t.Fatalf("%s != %s", Agnes.Animal(), s)
	}
}

func TestAgnesName(t *testing.T) {
	if ok := Agnes.Name() == agnes; !ok {
		t.Fatalf("%s != %s", Agnes.Name(), agnes)
	}
}
