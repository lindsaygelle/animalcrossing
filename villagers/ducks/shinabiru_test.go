package ducks

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestShinabiruAnimal(t *testing.T) {
	var s string = animals.Duck.Name()
	if ok := Shinabiru.Animal() == s; !ok {
		t.Fatalf("%s != %s", Shinabiru.Animal(), s)
	}
}

func TestShinabiruName(t *testing.T) {
	if ok := Shinabiru.Name() == shinabiru; !ok {
		t.Fatalf("%s != %s", Shinabiru.Name(), shinabiru)
	}
}
