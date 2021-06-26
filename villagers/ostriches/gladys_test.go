package ostriches

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestGladysAnimal(t *testing.T) {
	var s string = animals.Ostrich.Name()
	if ok := Gladys.Animal() == s; !ok {
		t.Fatalf("%s != %s", Gladys.Animal(), s)
	}
}

func TestGladysName(t *testing.T) {
	if ok := Gladys.Name() == gladys; !ok {
		t.Fatalf("%s != %s", Gladys.Name(), gladys)
	}
}
