package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTomNookAnimal(t *testing.T) {
	var s string = animals.Raccoon.Name()
	if ok := TomNook.Animal() == s; !ok {
		t.Fatalf("%s != %s", TomNook.Animal(), s)
	}
}

func TestTomNookName(t *testing.T) {
	if ok := TomNook.Name() == tomNook; !ok {
		t.Fatalf("%s != %s", TomNook.Name(), tomNook)
	}
}
