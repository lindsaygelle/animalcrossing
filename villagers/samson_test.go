package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestSamsonAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Samson.Animal() == s; !ok {
		t.Fatalf("%s != %s", Samson.Animal(), s)
	}
}

func TestSamsonName(t *testing.T) {
	if ok := Samson.Name() == samson; !ok {
		t.Fatalf("%s != %s", Samson.Name(), samson)
	}
}
