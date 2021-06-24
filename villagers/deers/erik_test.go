package deers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestErikAnimal(t *testing.T) {
	var s string = animals.Deer.Name()
	if ok := Erik.Animal() == s; !ok {
		t.Fatalf("%s != %s", Erik.Animal(), s)
	}
}

func TestErikName(t *testing.T) {
	if ok := Erik.Name() == erik; !ok {
		t.Fatalf("%s != %s", Erik.Name(), erik)
	}
}
