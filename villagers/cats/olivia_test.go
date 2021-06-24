package cats

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestOliviaAnimal(t *testing.T) {
	var s string = animals.Cat.Name()
	if ok := Olivia.Animal() == s; !ok {
		t.Fatalf("%s != %s", Olivia.Animal(), s)
	}
}

func TestOliviaName(t *testing.T) {
	if ok := Olivia.Name() == olivia; !ok {
		t.Fatalf("%s != %s", Olivia.Name(), olivia)
	}
}
