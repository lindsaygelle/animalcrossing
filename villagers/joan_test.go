package villagers

import (
	"testing"

	"github.com/lindsaygelle/animalcrossing/animals"
)

func TestJoanAnimal(t *testing.T) {
	var s string = animals.Boar.Name()
	if ok := Joan.Animal() == s; !ok {
		t.Fatalf("%s != %s", Joan.Animal(), s)
	}
}

func TestJoanName(t *testing.T) {
	if ok := Joan.Name() == joan; !ok {
		t.Fatalf("%s != %s", Joan.Name(), joan)
	}
}
