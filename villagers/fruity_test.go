package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestFruityAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Fruity.Animal() == s; !ok {
		t.Fatalf("%s != %s", Fruity.Animal(), s)
	}
}

func TestFruityName(t *testing.T) {
	if ok := Fruity.Name() == fruity; !ok {
		t.Fatalf("%s != %s", Fruity.Name(), fruity)
	}
}
