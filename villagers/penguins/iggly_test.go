package penguins

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestIgglyAnimal(t *testing.T) {
	var s string = animals.Penguin.Name()
	if ok := Iggly.Animal() == s; !ok {
		t.Fatalf("%s != %s", Iggly.Animal(), s)
	}
}

func TestIgglyName(t *testing.T) {
	if ok := Iggly.Name() == iggly; !ok {
		t.Fatalf("%s != %s", Iggly.Name(), iggly)
	}
}
