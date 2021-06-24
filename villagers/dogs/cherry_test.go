package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCherryAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Cherry.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cherry.Animal(), s)
	}
}

func TestCherryName(t *testing.T) {
	if ok := Cherry.Name() == cherry; !ok {
		t.Fatalf("%s != %s", Cherry.Name(), cherry)
	}
}
