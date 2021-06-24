package birds

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJacobAnimal(t *testing.T) {
	var s string = animals.Bird.Name()
	if ok := Jacob.Animal() == s; !ok {
		t.Fatalf("%s != %s", Jacob.Animal(), s)
	}
}

func TestJacobName(t *testing.T) {
	if ok := Jacob.Name() == jacob; !ok {
		t.Fatalf("%s != %s", Jacob.Name(), jacob)
	}
}
