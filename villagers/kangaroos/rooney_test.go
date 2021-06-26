package kangaroos

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRooneyAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Rooney.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rooney.Animal(), s)
	}
}

func TestRooneyName(t *testing.T) {
	if ok := Rooney.Name() == rooney; !ok {
		t.Fatalf("%s != %s", Rooney.Name(), rooney)
	}
}
