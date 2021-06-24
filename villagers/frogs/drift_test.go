package frogs

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestDriftAnimal(t *testing.T) {
	var s string = animals.Frog.Name()
	if ok := Drift.Animal() == s; !ok {
		t.Fatalf("%s != %s", Drift.Animal(), s)
	}
}

func TestDriftName(t *testing.T) {
	if ok := Drift.Name() == drift; !ok {
		t.Fatalf("%s != %s", Drift.Name(), drift)
	}
}
