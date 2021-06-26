package kangaroos

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCarrieAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Carrie.Animal() == s; !ok {
		t.Fatalf("%s != %s", Carrie.Animal(), s)
	}
}

func TestCarrieName(t *testing.T) {
	if ok := Carrie.Name() == carrie; !ok {
		t.Fatalf("%s != %s", Carrie.Name(), carrie)
	}
}
