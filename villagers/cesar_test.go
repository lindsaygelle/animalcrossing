package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestCesarAnimal(t *testing.T) {
	var s string = animals.Gorilla.Name()
	if ok := Cesar.Animal() == s; !ok {
		t.Fatalf("%s != %s", Cesar.Animal(), s)
	}
}

func TestCesarName(t *testing.T) {
	if ok := Cesar.Name() == cesar; !ok {
		t.Fatalf("%s != %s", Cesar.Name(), cesar)
	}
}
