package kangaroos

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSylviaAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Sylvia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Sylvia.Animal(), s)
	}
}

func TestSylviaName(t *testing.T) {
	if ok := Sylvia.Name() == sylvia; !ok {
		t.Fatalf("%s != %s", Sylvia.Name(), sylvia)
	}
}
