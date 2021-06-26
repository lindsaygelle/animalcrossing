package sheep

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestVestaAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Vesta.Animal() == s; !ok {
		t.Fatalf("%s != %s", Vesta.Animal(), s)
	}
}

func TestVestaName(t *testing.T) {
	if ok := Vesta.Name() == vesta; !ok {
		t.Fatalf("%s != %s", Vesta.Name(), vesta)
	}
}
