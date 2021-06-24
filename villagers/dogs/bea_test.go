package dogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBeaAnimal(t *testing.T) {
	var s string = animals.Dog.Name()
	if ok := Bea.Animal() == s; !ok {
		t.Fatalf("%s != %s", Bea.Animal(), s)
	}
}

func TestBeaName(t *testing.T) {
	if ok := Bea.Name() == bea; !ok {
		t.Fatalf("%s != %s", Bea.Name(), bea)
	}
}
