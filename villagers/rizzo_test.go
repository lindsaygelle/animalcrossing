package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestRizzoAnimal(t *testing.T) {
	var s string = animals.Mouse.Name()
	if ok := Rizzo.Animal() == s; !ok {
		t.Fatalf("%s != %s", Rizzo.Animal(), s)
	}
}

func TestRizzoName(t *testing.T) {
	if ok := Rizzo.Name() == rizzo; !ok {
		t.Fatalf("%s != %s", Rizzo.Name(), rizzo)
	}
}
