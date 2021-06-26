package sheep

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTimbraAnimal(t *testing.T) {
	var s string = animals.Sheep.Name()
	if ok := Timbra.Animal() == s; !ok {
		t.Fatalf("%s != %s", Timbra.Animal(), s)
	}
}

func TestTimbraName(t *testing.T) {
	if ok := Timbra.Name() == timbra; !ok {
		t.Fatalf("%s != %s", Timbra.Name(), timbra)
	}
}
