package rhinoceroses

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestTiaraAnimal(t *testing.T) {
	var s string = animals.Rhinoceros.Name()
	if ok := Tiara.Animal() == s; !ok {
		t.Fatalf("%s != %s", Tiara.Animal(), s)
	}
}

func TestTiaraName(t *testing.T) {
	if ok := Tiara.Name() == tiara; !ok {
		t.Fatalf("%s != %s", Tiara.Name(), tiara)
	}
}
