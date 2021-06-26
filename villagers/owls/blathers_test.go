package owls

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestBlathersAnimal(t *testing.T) {
	var s string = animals.Owl.Name()
	if ok := Blathers.Animal() == s; !ok {
		t.Fatalf("%s != %s", Blathers.Animal(), s)
	}
}

func TestBlathersName(t *testing.T) {
	if ok := Blathers.Name() == blathers; !ok {
		t.Fatalf("%s != %s", Blathers.Name(), blathers)
	}
}
