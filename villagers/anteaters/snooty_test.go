package anteaters

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSnootyAnimal(t *testing.T) {
	var s string = animals.Anteater.Name()
	if ok := Snooty.Animal() == s; !ok {
		t.Fatalf("%s != %s", Snooty.Animal(), s)
	}
}

func TestSnootyName(t *testing.T) {
	if ok := Snooty.Name() == snooty; !ok {
		t.Fatalf("%s != %s", Snooty.Name(), snooty)
	}
}
