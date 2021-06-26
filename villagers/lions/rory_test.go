package lions

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRoryAnimal(t *testing.T) {
	var s string = animals.Lion.Name()
	if ok := Rory.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rory.Animal(), s)
	}
}

func TestRoryName(t *testing.T) {
	if ok := Rory.Name() == rory; !ok {
		t.Fatalf("%s != %s", Rory.Name(), rory)
	}
}
