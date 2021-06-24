package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestEmeraldAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Emerald.Animal() == s; !ok {
		t.Fatalf("%s != %s", Emerald.Animal(), s)
	}
}

func TestEmeraldName(t *testing.T) {
	if ok := Emerald.Name() == emerald; !ok {
		t.Fatalf("%s != %s", Emerald.Name(), emerald)
	}
}
