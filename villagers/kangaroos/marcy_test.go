package kangaroos

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestMarcyAnimal(t *testing.T) {
	var s string = animals.Kangaroo.Name()
	if ok := Marcy.Animal() == s; !ok {
		t.Fatalf("%s != %s", Marcy.Animal(), s)
	}
}

func TestMarcyName(t *testing.T) {
	if ok := Marcy.Name() == marcy; !ok {
		t.Fatalf("%s != %s", Marcy.Name(), marcy)
	}
}
