package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestPortiaAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Portia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Portia.Animal(), s)
	}
}

func TestPortiaName(t *testing.T) {
	if ok := Portia.Name() == portia; !ok {
		t.Fatalf("%s != %s", Portia.Name(), portia)
	}
}
