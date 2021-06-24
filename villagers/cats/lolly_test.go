package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestLollyAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Lolly.Animal() == s; !ok {
		t.Fatalf("%s != %s", Lolly.Animal(), s)
	}
}

func TestLollyName(t *testing.T) {
	if ok := Lolly.Name() == lolly; !ok {
		t.Fatalf("%s != %s", Lolly.Name(), lolly)
	}
}
