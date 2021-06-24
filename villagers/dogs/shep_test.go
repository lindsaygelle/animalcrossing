package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestShepAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Shep.Animal() == s; !ok {
		t.Fatalf("%s != %s", Shep.Animal(), s)
	}
}

func TestShepName(t *testing.T) {
	if ok := Shep.Name() == shep; !ok {
		t.Fatalf("%s != %s", Shep.Name(), shep)
	}
}
